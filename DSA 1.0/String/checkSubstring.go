package main

import "strings"

func checksubstring(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}
