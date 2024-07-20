package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Is Hello is Palindrome", checkPalindrom("hello"))
	fmt.Println("Is Madam is Palindrome", checkPalindrom("Madam"))

	str := "Shahabaz Sulthan"
	fmt.Println(str)
	field := strings.Fields(str)
	fmt.Println(field)

	fmt.Println("Reverse Of Hello World = ", reverseWords("Hello World"))
	fmt.Println("Reverse Of Shahabaz Sulthan = ", reverseWords("Shahabaz Sulthan"))

	Char := countEachChar("Shahabaz")
	for c, v := range Char {
		fmt.Printf("%c -> %v\n", c, v)
	}

	fmt.Println("Revese Of Shahabaz = ", reverseChar("shahabaz"))

	fmt.Println("First unique character index of leetcode = ", firstUniqueChar("leetcode"))

	var strs = []string{"flower","flow","flood"}
	fmt.Println("Long Common Prefix = ",LongCommonPrefix(strs))

	fmt.Println("remove dupilcate From 'programming' = ",removeDuplicate("Programming"))

	s := "Hello, World!"
    vowels, consonants := countVowelsandConstants(s)
    fmt.Printf("Vowels: %d, Consonants: %d\n", vowels, consonants)

	fmt.Println("Hello world is hello contains = ",checksubstring("Hello World","World"))

	ans := `this is a test this is only a test`
    counts := wordFrequency(ans)
    for word, count := range counts {
        fmt.Printf("%s: %d\n", word, count)
    }

	str1 := "remove all spaces from this string"
    fmt.Println(removeSpace(str1))

	input := "hello"
    reversed := reveseString(input)
    fmt.Println("Original string:", input)
    fmt.Println("Reversed string:", reversed)
}
