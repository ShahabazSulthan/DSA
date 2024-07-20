package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans = append(ans, i, j)
			}
		}
	}
	return ans
}

func main() {
	var arr1 = []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr1, 9))
	// time complexity 0
	fmt.Println("Time Reduced",SumOfTwo(arr1,13)) 

}
