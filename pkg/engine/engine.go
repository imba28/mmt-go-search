package engine

import (
	"log"
	"time"
)

// Result ... represents a search result
type Result struct {
	Type        string
	Title       string
	URL         string
	Description string
}

var (
	webSearch      = fakeSearch("Web")
	imageSearch    = fakeSearch("Image")
	videoSearch    = fakeSearch("Video")
	newsSearch     = fakeSearch("News")
	shoppingSearch = fakeSearch("Shopping")
)

// firstResult ... sends requests to all replicas and returns only the fastest one
func firstResult(query string, replicas ...search) []Result {
	// todo
}

// LoadResults - sends requests to different search backends and aggregates the result.
func LoadResults(query string) []Result {
	var results []Result
	backends := []search{
		webSearch,
		imageSearch,
		videoSearch,
		newsSearch,
		shoppingSearch,
	}
	c := make(chan []Result)

	/*
		todo: Add multiple replicas for every search backend.
		Send requests to all of them and use only the fastest result.

		You will need:
		- func, closures
		- chan []Result
		- Goroutines

		hint: Fan-In pattern
		https://github.com/imba28/mmt-go-snippets/blob/master/3-concurrency/3-select/main.go#L17
	*/

	for i := 0; i < len(backends); i++ {
		go func(backend search) {
			c <- backend(query)
		}(backends[i])
	}

	timeout := time.After(300 * time.Millisecond)
	for i := 0; i < len(backends); i++ {
		select {
		case r := <-c:
			results = append(results, r...)
		case <-timeout:
			log.Println("search timeout! Skipping results")
		}

	}

	return results
}
