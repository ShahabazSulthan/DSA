package main

import "fmt"

func SumOfTwo(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		ans := target - num
		if j, ok := numMap[ans]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return []int{}
}

func main() {
	var arr1 = []int{2, 7, 11, 15}
	fmt.Println(SumOfTwo(arr1, 13))

}
