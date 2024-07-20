package main

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[: n-1]
	return x
}

func heapSort(arr []int) []int {
	h := &maxHeap{}
	heap.Init(h)

	for _, i := range arr {
		heap.Push(h, i)
	}

	sorted := make([]int, 0, len(*h))

	for len(*h) > 0 {
		sorted = append(sorted, heap.Pop(h).(int))
	}

	return sorted
}

func main() {
	h := &maxHeap{2, 1, 5}
	heap.Init(h)

	heap.Push(h, 3)

	fmt.Printf("max : %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}

	arr1 := []int{8,2,5,1,8,3,1}

	fmt.Println("Sorted Heap : ",heapSort(arr1))
}
