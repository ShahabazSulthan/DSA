package main

func sumOfDigit(n int) int {
	if n == 0 {
		return 0
	}

	return n % 10 + sumOfDigit(n/10)
}
