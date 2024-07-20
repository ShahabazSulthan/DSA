package main

func findMin(arr []int) int {

	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left]
}

func findMax(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		// We need to find the peak element which is greater than its neighbors
		if arr[mid] > arr[right] {
			// max is in the right half
			left = mid + 1
		} else {
			// max is in the left half including mid
			right = mid
		}
	}
	// The maximum element is at the position just before the smallest element found
	// Therefore, we should return arr[left - 1]
	if left == 0 {
		return arr[len(arr)-1]
	}
	return arr[left-1]

}

