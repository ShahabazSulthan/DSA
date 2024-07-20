package main

import "fmt"

func main() {

	res := factorial(5)
	fmt.Println(res) // Output: 120

	result1 := fibonacci(10)
    fmt.Println(result1) // Output: 55

	number := 12345
	result := sumOfDigit(number)
	fmt.Printf("Sum of digits of %d is %d\n", number, result)

	fmt.Println("Hello revesed = ",reverse("hello"))

	fmt.Println("madam is a palidrom",isPalindrome("madam"))

	fmt.Println(power(2,3))


}
