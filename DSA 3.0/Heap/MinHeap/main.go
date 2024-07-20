package main

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sortHeap(arr []int) []int {
	h := &minHeap{}
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
	h := &minHeap{2, 1, 5}
	heap.Init(h)

	heap.Push(h, 3)

	fmt.Printf("min : %d\n", (*h)[0])

	for len(*h) > 0 {
		fmt.Println(heap.Pop(h))
	}

	arr1 := []int{8, 1, 4, 2, 9, 3}

	fmt.Println("Sorted Heap : ", sortHeap(arr1))

}
