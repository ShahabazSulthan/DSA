package main

import "fmt"

type Queue struct {
	element []int
}

func (q *Queue) Enqueue(val int) {
	q.element = append(q.element, val)
}

func (q *Queue) Dequeue() int {
	if len(q.element) == 0 {
		fmt.Println("Queue is empty")
		return 0
	}

	Front := q.element[0]
	q.element = q.element[1:]
	return Front
}

func deleteSpecQueue(q *Queue,value int) {
	if len(q.element) == 0 {
		fmt.Println("Queue is empty")
        return
	}

	top := q.Dequeue()
	if top == value {
		return
	}

	deleteSpecQueue(q,value)

	q.Enqueue(top)
}

func main() {
	q := &Queue{}

	arr := []int{9,7,5,3,1}

	for _,v := range arr {
		q.Enqueue(v)
	}

	fmt.Println("Before Delete = ",q.element)

	deleteSpecQueue(q,5)

	fmt.Println("After delete = ",q.element)
}
