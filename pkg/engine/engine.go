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

	results = append(results, webSearch(query)...)
	results = append(results, imageSearch(query)...)
	results = append(results, videoSearch(query)...)
	results = append(results, newsSearch(query)...)
	results = append(results, shoppingSearch(query)...)

	return results
}
