package engine

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

// LoadResults - sends requests to different search backends and aggregates the result.
func LoadResults(query string) []Result {
	var results []Result

	/*
		todo: fakeSearch(...) calls should be executed concurrently
		you will need:
			 - chan []Result
			 - goroutines
			 - closure or func
	*/

	backends := []search{
		webSearch,
		imageSearch,
		videoSearch,
		newsSearch,
		shoppingSearch,
	}

	for i := range backends {
		results = append(results, backends[i](query)...)
	}

	return results
}
