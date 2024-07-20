package main

func reverse(s string) string {
	if len(s) == 0 {
		return s
	}

	return reverse(s[1:]) + string(s[0])
}
