package api

import (
	"fmt"
	"html/template"
	"log"
	"mmt/go-search/pkg/engine"
	"net/http"
	"time"
)

type searchView struct {
	Query       string
	TimeElapsed time.Duration
	Results     []engine.Result
}

// result page
func resultsPageHandler() http.HandlerFunc {
	t := initTemplate(template.New("search"), nil)

	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")

		start := time.Now()
		searchResults := engine.LoadResults(q)
		timeElapsed := time.Since(start).Round(time.Millisecond)

		model := searchView{
			Query:       q,
			TimeElapsed: timeElapsed,
			Results:     searchResults,
		}

		if err := t.Execute(w, model); err != nil {
			log.Println("Could not execute search template!")
		}
		log.Println(fmt.Sprintf("%d results, took %v", len(searchResults), timeElapsed))
	}
}

// home page
func homePageHandler() http.HandlerFunc {
	t := initTemplate(template.New("index"), nil)

	return func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, nil); err != nil {
			log.Println("Could not execute search template!")
		}
	}
}
