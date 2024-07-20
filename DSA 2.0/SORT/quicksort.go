package main

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}

	left,right := 0,len(a)-1

	pivot := len(a)/2

	a[pivot],a[right] = a[right],a[pivot]
	
	for i := range a {
		if a[i] < a[right] {
			a[i],a[left] = a[left],a[i]
			left++
		}
	}

	a[left],a[right] = a[right],a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])
}