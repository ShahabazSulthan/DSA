package main

func FindMin(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	min := arr[0]
	for _, val := range arr[1:] {
		if val < min {
			min = val
		}
	}
	return min
}

func FindMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	max := arr[0]
	for _, val := range arr[1:] {
		if val > max {
			max = val
		}
	}
	return max
}
