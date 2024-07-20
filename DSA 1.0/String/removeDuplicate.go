package main

func removeDuplicate(str string) string {
	Maps := make(map[rune]bool)
	result := []rune{}

	for _,c := range str {
		if !Maps[c] {
			Maps[c]=true
			result = append(result, c)
		}
	}
	return string(result)
}
