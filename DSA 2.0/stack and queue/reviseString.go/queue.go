package main

import "fmt"

type Queue struct {
	element []rune
}

func (q *Queue) Enqueue(input rune) {
	q.element = append(q.element, input)
}

func (q *Queue) isEmpty() bool {
	return len(q.element) == 0
}

func (q *Queue) Dequeue() rune {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
	}

	top := q.element[0]
	q.element = q.element[1:]
	return top
}

func reviseString(input string) string {
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

func main() {
	input := "Hello everybody"
	fmt.Println("revise = ",reviseString(input))
}
