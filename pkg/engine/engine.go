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
			c <- backend(query)
		}(backends[i])
	}

	for i := 0; i < len(backends); i++ {
		results = append(results, <-c...)
	}

	return results
}
