package main

import (
	"strings"
	"unicode"
)

func countVowelsandConstants(str string) (int, int) {
	vomels := "aeiouAEIOU"
	vowelsCount := 0
	connstants := 0

	for _, c := range str {
		if unicode.IsLetter(c) {
			if strings.ContainsRune(vomels, c) {
				vowelsCount++
			} else {
				connstants++
			}
		}
	}
	return vowelsCount, connstants
}
