package main

func merge(left,right []int) []int {
	res := make([]int,0,len(left)+len(right))

	i,j :=0,0

	for i < len(left) && j < len(right) {
		if left[i] < right [j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	for i < len(left) {
		res = append(res, left[i])
		i++
	}

	for j < len(right) {
		res = append(res, right[j])
		j++
	}

	return res
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	mid := len(a)/2
	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])
	return merge(left,right)
}