// Mervan Kaya palinda21 uppg 2.3
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	listofwords := strings.Fields(s) // Use fields to make list of words
	m := make(map[string]int)        // Create the map
	for _, v := range listofwords {  // Iterate over each word
		i := 0
		for _, val := range listofwords { // Count how many occurences that word has
			if v == val {
				i += 1
			}
		}
		m[v] = i // Add the amount of occurences to the map along with key
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
