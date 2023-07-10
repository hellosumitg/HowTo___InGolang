package main

import (
	"log"
)

func main() {
	// Create a new instance of CatFactService with the provided URL.
	svc := NewCatFactService("https://catfact.ninja/fact")

	// Wrap the service with LoggingService to log execution time and errors.
	svc = NewLoggingService(svc)

	// fact, err := svc.GetCatFact(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", fact)

	// Create a new instance of ApiServer with the wrapped service.
	apiServer := NewApiServer(svc)

	// Start the API server and log any errors.
	log.Fatal(apiServer.Start(":3000"))
}
