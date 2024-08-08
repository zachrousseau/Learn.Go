package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	var words []string = strings.Fields(s)

	var word_map = make(map[string]int)

	for _, word := range words { 
		word_map[word] = word_map[word] + 1
	}

	return word_map
}

func main() {
	wc.Test(WordCount)
}
