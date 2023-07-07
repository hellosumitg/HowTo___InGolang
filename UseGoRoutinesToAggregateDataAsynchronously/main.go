package main

import (
	"fmt"
	"sync"
	"time"
)

// how to aggregate data in an asynchronous Way by using golang's concurrency model,
// so we can cut down a big portion of our HTTP round trip and make out applications faster

func main() {
	start := time.Now()
	userName := fetchUser() // 100ms

	// ---------------------------------`Slower Method of Fetching Data`-----------------------------------//
	// likes := fetchUserLikes(userName)
	// match := fetchUserMatch(userName)
	// fmt.Println("likes: ", likes)
	// fmt.Println("match: ", match)
	// fmt.Println("Data Fetching/Aggregation Non-Asynchronously takes: ", time.Since(start))
	// ---------------------------------`Slower Method of Fetching Data`-----------------------------------//

	// ---------------------------------`Faster Method of Fetching Data Asynchronously`-----------------------------------//
	response := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fetchUserLikes(userName, response, wg)
	go fetchUserMatch(userName, response, wg)

	wg.Wait() // block until 2 wg.Done()
	close(response)

	for resp := range response {
		fmt.Println("response: ", resp)
	}

	fmt.Println("Data Fetching/Aggregation Asynchronously takes: ", time.Since(start))
	// ---------------------------------`Faster Method of Fetching Data Asynchronously`-----------------------------------//
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

// ---------------------------------`Slower Method of Fetching Data`-----------------------------------//

// func fetchUserLikes(userName string) int {
// 	time.Sleep(time.Millisecond * 150)

// 	return 11
// }

// func fetchUserMatch(userName string) string {
// 	time.Sleep(time.Millisecond * 100)

// 	return "ANNA"
// }

// ---------------------------------`Slower Method of Fetching Data`-----------------------------------//

// ---------------------------------`Faster Method of Fetching Data Asynchronously`-----------------------------------//

func fetchUserLikes(userName string, response chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	response <- 11
	wg.Done()
}

func fetchUserMatch(userName string, response chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	response <- "ANNA"
	wg.Done()
}

// ---------------------------------`Faster Method of Fetching Data Asynchronously`-----------------------------------//
