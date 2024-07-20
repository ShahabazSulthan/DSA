package main

func reveseString(str string) string {
	if len(str) <= 1 {
		return str
	}

	return string(str[len(str)-1]) + reveseString(str[:len(str)-1])
}
