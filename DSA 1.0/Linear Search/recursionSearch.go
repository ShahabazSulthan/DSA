package main

func recursionSearch(arr []int, target int, index int) int {
	if index >= len(arr) {
		return -1
	}

	if arr[index] == target {
		return index
	}

	index++
	return recursionSearch(arr, target, index)

}
