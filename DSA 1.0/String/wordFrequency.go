package main

import "strings"

func wordFrequency(str string) map[string]int {
	seen := make(map[string]int)
	words := strings.Fields(str)
	for _,word := range words {
		seen[word]++
	}
	return seen
}