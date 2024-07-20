package main

func seletionSort(a []int) []int {
	n := len(a)
	for i:=0;i<n-1;i++ {
		index := i
		for j:=i+1;j<n;j++ {
			if a[j] < a[index] {
				index = j
			}
		}
		a[i],a[index] = a[index],a[i]
	}
	return a
}