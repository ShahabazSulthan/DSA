package main

import "strings"

func removeSpace(str string) string {
	return strings.ReplaceAll(str," ","")
}