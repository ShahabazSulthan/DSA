package main

import "fmt"

type Queue struct {
	elements []int
}

func (q *Queue) isEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Enqueue(val int) {
	q.elements = append(q.elements, val)
}

func (q *Queue) Dequeue() int {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return 0
	}

	Front := q.elements[0]
	q.elements = q.elements[1:]
	return Front
}

func (q *Queue) Peek() int {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return 0
	}

	peek := q.elements[0]
	return peek
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func main() {
	q := Queue{}

	q.Enqueue(20)
	q.Enqueue(40)
	q.Enqueue(60)
	q.Enqueue(80)

	fmt.Println("Queue = ",q.elements)

	fmt.Println("Size of queue = ",q.Size())

	fmt.Println("Front of Dequeued queue = ",q.Dequeue())

	fmt.Println("Queue = ",q.elements)

	fmt.Println("Peek = ",q.Peek())

	fmt.Println("is Queue empty = ",q.isEmpty())
}