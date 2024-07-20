package main

func bubbleSort(a []int) []int {
	n := len(a)
	for i:=0;i<n-1;i++ {
		for j:=0;j<n-1;j++ {
		   if a[j] > a[j+1] {
			a[j],a[j+1]=a[j+1],a[j]
		   }
		}
	}
	return a
}