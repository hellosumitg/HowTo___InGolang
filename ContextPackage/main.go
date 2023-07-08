package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	// ctx := context.Background() // Uncomment it first time
	ctx := context.WithValue(context.Background(), "foo", "bar") // Create a new context with a key-value pair, where "foo" is the key and "bar" is the value
	userID := 10
	val, err := fetchUserData(ctx, userID) // Fetch user data using the provided context
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: ", val)
	fmt.Println("It Took: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

/*
func fetchUserData(ctx context.Context, userID int) (int, error) {
	val, err := fetchThirdPartyStuffWhichCanBeSlow()
	if err != nil {
		return 0, err
	}

	return val, nil
}
*/

/*
fetchUserData is a function that fetches user data using the provided context and user ID.
It encapsulates the process of calling a potentially slow third-party operation while respecting the context's timeout.
The function returns the fetched value and any error that occurred during the process.
*/
func fetchUserData(ctx context.Context, userID int) (int, error) {
	val := ctx.Value("foo") // Retrieve the value associated with the key "foo" from the context
	fmt.Println(val)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200) // Create a new context with a timeout of 200 milliseconds
	defer cancel()
	resChan := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow() // Perform a potentially slow operation
		resChan <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done(): // If the context times out
			return 0, fmt.Errorf("Fetching data from the third party took too long")
		case res := <-resChan: // If the slow operation completes
			return res.value, res.err
		}
	}
}

/*
fetchThirdPartyStuffWhichCanBeSlow simulates a slow operation by sleeping for 150 milliseconds.
In a real-world scenario, this function could represent an actual interaction with a third-party service.
It returns a value and any error that occurred during the process.
*/
func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	// time.Sleep(time.Millisecond * 500) // first time
	time.Sleep(time.Millisecond * 150) // Simulate a slow operation // second time
	return 666, nil
}

/*
Practical Application of the above example in a real-world scenario:-
You can utilize the `context` package in Golang to efficiently manage shared state among multiple concurrent goroutines.
By encapsulating the shared state within a `context` value, you can achieve interesting benefits.
Let's consider a practical use case involving the use of `request IDs`.
Suppose you have a middleware system in your microservices architecture.
Whenever a request arrives, you can generate a unique `request ID` and store it within the `context`.
As your application proceeds to invoke various functions and services, this `request ID` remains accessible throughout the execution flow.
In the event of an error occurring within any of the functions, you can easily associate it with the corresponding `request ID`.
This allows you to effectively trace user behavior and track the application's performance using tools like Grafana or ElasticSearch.
By incorporating the `request ID` into the `context`, you establish a reliable mechanism for tracing and monitoring your application.
This approach enhances observability, making it easier to identify and troubleshoot issues.
*/
