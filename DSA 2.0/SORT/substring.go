package main

import "fmt"

func FindlongestsubString(s string) int {
	mapp := make(map[byte]int)
	start := 0
	maxLength := 0

	for end := 0; end < len(s); end++ {
		if index, found := mapp[s[end]]; found && index >= start {
			start = index + 1
		}
		mapp[s[end]] = end
		if end-start > maxLength {
			maxLength = end - start + 1
		}
	}
	return maxLength
}

func main() {
	input := "abcabbcabc"

	fmt.Printf("The length of the longest substring without repeating characters in \"%s\" is %d\n", input, FindlongestsubString(input))
}
