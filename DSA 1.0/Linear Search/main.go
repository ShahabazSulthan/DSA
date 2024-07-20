package main

import "fmt"

func main() {
	var arr = []int{6, 2, 1, 8, 3, 1}
	target := 3
	index := findElementpos(arr, target)
	fmt.Printf("Element %d found at index %d\n", target, index)

	var arr2 = []int{5, 3, 1, 5, 8, 5, 2}
	target2 := 5
	count := HowmanyTimeOccur(arr2, target2)
	fmt.Printf("Element %d occurs %d times in the list \n", target2, count)

	Max := FindMax(arr)
	Min := FindMin(arr)

	fmt.Printf("Minimum value in the %d is %d\n", arr, Min)
	fmt.Printf("Maximum value in the %d is %d\n", arr, Max)

	var arr3 = []int{9, 4, 7, 5, 3, 6}
	target3 := 7
	index1 := recursionSearch(arr3, target3, 0)
	if index1 != -1 {
		fmt.Printf("Element %d found at index %d\n", target3, index1)
	} else {
		fmt.Printf("Element %d not found in the array\n", target3)
	}
}
