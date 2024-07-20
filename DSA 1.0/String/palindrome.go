package main

import "strings"

func checkPalindrom(str string) bool {
	str = strings.ToLower(str)
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}
