package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	Rear  *Node
	Front *Node
	size  int
}

func (q *Queue) Enqueue(val int) {
	newNode := &Node{data: val}
	if q.Rear != nil {
		q.Rear.next = newNode
	}
	q.Rear = newNode
	if q.Front == nil {
		q.Front = newNode
	}
	q.size++
}

func (q *Queue) Pop() int {
	if q.Front == nil {
		fmt.Println("queue is empty")
		return 0
	}

	front := q.Front.data
	q.Front = q.Front.next
	if q.Front == nil {
		q.Rear = nil
	}
	q.size--
	return front
}

func (q *Queue) Peek() int {
	if q.Front == nil {
		fmt.Println("Queue is empty")
		return 0
	}

	return q.Front.data
}

func (q *Queue) isEmpty() bool {
	return q.Front == nil
}


func (q *Queue) Print() {
	if q.Front == nil {
		fmt.Println("Queue is empty")
		return 
	}

	current := q.Front
	for current != nil {
		fmt.Printf("%d -> ",current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	q := Queue{}

	q.Enqueue(200)
	q.Enqueue(400)
	q.Enqueue(600)
	q.Enqueue(800)

	q.Print()

	fmt.Println("Peek = ",q.Peek())

	fmt.Println("Size Of Queue",q.size)

	fmt.Println("Dequeued Front element = ",q.Pop())

	q.Print()

	fmt.Println("Peek = ",q.Peek())

	fmt.Println("Is Queue Empty = ",q.isEmpty())
}