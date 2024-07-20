package main

import "fmt"

func main() {
	arr := []int{34, 64, 12, 86, 56, 98}

	fmt.Println("Bubble Sort = ",bubbleSort(arr))

	arr1 := []int{8,3,1,6,4,9}

	fmt.Println("Insertion Sort = ",insertionSort(arr1))

	fmt.Println("Selection Sort = ",seletionSort(arr))

	arr3 := []int{500,200,600,400,900}

	quickSort(arr3)

	fmt.Println("Quick Sort = ",arr3)

	fmt.Println("Merge Sort = ",mergeSort(arr))
}
