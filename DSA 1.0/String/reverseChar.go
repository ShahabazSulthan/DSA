package main

func reverseChar(str string) string {
	runes := []rune(str)
	n := len(str)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
