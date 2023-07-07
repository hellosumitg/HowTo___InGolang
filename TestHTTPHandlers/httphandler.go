package httphandler

import "net/http"

func handleGet(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Uncomment the following line to set a custom header "x request id"
	// w.Header().Set("x request id", "dsfsdfjsefi")

	w.WriteHeader(http.StatusOK) // Set the response status code to 200

	// Uncomment the following line to test a different status code
	// w.WriteHeader(http.StatusBadGateway) // When running the test, it will output: `httphandler_test.go:21: Expected 200 but got 502`

	w.Write([]byte("GOOD")) // Write the response body as "GOOD"

	// Uncomment the following line to test a different response body
	// w.Write([]byte("BAD")) // When running the test, it will output: `httphandler_test.go:32: Expected GOOD but got BAD`
}
