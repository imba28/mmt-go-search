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

func firstResult(query string, replicas ...search) []Result {
	c := make(chan []Result)
	for i := range replicas {
		go func(i int) {
			c <- replicas[i](query)
		}(i)
	}

	return <-c
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

	for i := 0; i < len(backends); i++ {
		go func(backend search) {
			c <- firstResult(query, backend, backend, backend)
		}(backends[i])
	}

	timeout := time.After(300 * time.Millisecond)
	for i := 0; i < len(backends); i++ {
		select {
		case r := <-c:
			results = append(results, r...)
		case <-timeout:
			log.Println("search timeout! Skipping results.")
			return results
		}

	}

	return results
}
