package main

import "fmt"

type node struct {
	data rune
	next *node
}

type Queue struct {
	Front *node
	Rear  *node
}

func (q *Queue) Enqueue(val rune) {
	newNode := &node{data: val}
	if q.Rear != nil {
		q.Rear.next = newNode
	}
	q.Rear = newNode
	if q.Front == nil {
		q.Front = newNode
	}
}

func (q *Queue) Dequeue() rune {
	if q.Front == nil {
		fmt.Println("Queue is empty")
	}

	Front := q.Front.data
	q.Front = q.Front.next
	return Front
}

func (q *Queue) isEmpty() bool {
	return q.Front == nil
}

func reviseQueue(input string) string {
	q := Queue{}

	for _,char := range input {
		q.Enqueue(char)
	}

	var revise string

	for !q.isEmpty() {
		revise = string(q.Dequeue()) + revise
	}

	return revise
}

