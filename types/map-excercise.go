package types

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

/*
WordCount - Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the
provided function and prints success or failure.
*/
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

// MapExcercise - Invoke WordCount() method
func MapExcercise() {
	fmt.Printf("****Running types.MapExcercise(), running go tour map excercise \n")
	wc.Test(WordCount)
}
