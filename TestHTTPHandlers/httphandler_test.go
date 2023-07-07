package httphandler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Actually there are two ways that we can use for testing, provided by the Go standard library:
1) HTTP Test Server
2) HTTP Response Recorder
Both methods serve a similar purpose, but they have some differences. Let's examine them one by one...
*/

// 1) HTTP Test Server
func TestHandleGetTestServer(t *testing.T) {
	// Create a new test server with the handleGet function as the handler
	server := httptest.NewServer(http.HandlerFunc(handleGet))

	// Send a GET request to the test server
	response, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but got %d", response.StatusCode)
	}

	defer response.Body.Close()

	// Check the response body content
	expected := "GOOD"
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("Expected %s but got %s", expected, string(b))
	}
}

// 2) HTTP Response Recorder
func TestHandleGetResponseRecorder(t *testing.T) {
	// Create a new response recorder
	responseRecorder := httptest.NewRecorder()

	/*
		The HTTP Response Recorder serves a similar purpose to an "HTTP Test Server" by capturing all content written with the response writer.
		It enables testing of different aspects, including header examination.
		However, identifying the origin of a specific header can be difficult.
		It may have been set by our server using code such as `w.Header().Set("x request id", "dsfsdfjsefi")`,
		or it could have originated from a proxy, reverse proxy, or another source.
		The HTTP Response Recorder is valuable in such scenarios as it helps trace the source of headers and facilitates testing of reverse proxies and related components.
	*/

	// Create a new HTTP request
	request, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}

	// Call the handleGet function with the response recorder and request
	handleGet(responseRecorder, request)

	// Check the response status code
	if responseRecorder.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but got %d", responseRecorder.Result().StatusCode)
	}

	defer responseRecorder.Result().Body.Close()

	// Check the response body content
	expected := "GOOD"
	b, err := ioutil.ReadAll(responseRecorder.Result().Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("Expected %s but got %s", expected, string(b))
	}
}

// Command for running the tests: `go test -v ./...` The "-v" flag is for verbose output, which prints all the log lines present in a test.

/*
In another scenario, instead of dealing with an HTTP response or request, we might be working with JSON data.
However, the underlying principle remains the same. Rather than using `ioutil.ReadAll` to read the data,
we can utilize a `JsonDecoder` to decode the JSON into a structured format.
This allows us to verify if the values retrieved from our database are correct, ensuring the integrity of the data.
*/
