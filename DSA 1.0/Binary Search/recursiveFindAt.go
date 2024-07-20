package main

func recursiveFind(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return recursiveFind(arr, target, mid+1, right)
	} else {
		return recursiveFind(arr, target, left, mid-1)
	}
}
