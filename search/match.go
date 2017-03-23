package search

import (
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(searchTerm string, feed *Feed) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(searchTerm, feed)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range Result {
		results <- result
	}
}

func Display(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
