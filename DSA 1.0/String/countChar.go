package main

func countEachChar(str string) map[rune]int {

	mapp := make(map[rune]int)

	for _, v := range str {
		mapp[v]++
	}
	return mapp
}
