package main

func firstUniqueChar(str string) int {
	countChar := make(map[rune]int)

	for _, char := range str {
		countChar[char]++
	}

	for i, c := range str {
		if countChar[c] == 1 {
			return i
		}
	}
	return -1
}
