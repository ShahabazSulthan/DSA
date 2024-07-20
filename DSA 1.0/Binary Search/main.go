package main

import "fmt"

func main() {
	var arr1 = []int{44, 66, 77, 88, 99}
	target := 66
	index := FindElementAt(arr1, target)
	fmt.Printf("Element %d found at index %d\n", target, index)

	var arr2 = []int{20, 30, 40, 50, 60}
	target2 := 50
	index2 := recursiveFind(arr2, target2, 0, len(arr2)-1)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target2, index2)
	} else {
		fmt.Printf("Element %d not found in the array\n", target2)
	}

	var arr3 = []int{10, 20, 20, 20, 30, 40, 50}
	target3 := 20
	index3 := FirstOcuurence(arr3, target3)
	fmt.Printf("First occurrence of element %d found at index %d \n", target3, index3)

	index4 := LastOcuurence(arr3, target3)
	fmt.Printf("Last occurrence of element %d found at index %d \n", target3, index4)

	var arr4 = []int{40, 20, 30, 50, 10}
	min := findMin(arr4)
	max := findMax(arr4)
	fmt.Println("Min = ",min, " Max = ",max)
}
