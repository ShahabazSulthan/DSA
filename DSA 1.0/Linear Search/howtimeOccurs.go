package main

func HowmanyTimeOccur(arr []int,target int) int {
	count := 0
	for _,num := range arr {
		if num == target {
			count++
		}
	}
	return count
}