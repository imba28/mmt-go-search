package engine

import (
	"fmt"
	"math/rand"
	"time"
)

type search func(query string) []Result

// creates a closure to a fake search backend service.
func fakeSearch(resultType string) search {
	return func(query string) []Result {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		return randomResults(resultType, 5)
	}
}

func randomResults(resultType string, n int) []Result {
	var results []Result
	for i := 0; i < rand.Intn(5)+1; i++ {
		results = append(results, Result{
			Type:        resultType,
			Title:       fmt.Sprintf("%s Result %d", resultType, i),
			URL:         "http://example.com/",
			Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.",
		})
	}
	return results
}
