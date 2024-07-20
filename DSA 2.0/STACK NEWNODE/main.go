package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type stack struct {
	top *Node
}

func (s *stack) Push(v int) {
	newNode := &Node{data: v, next: s.top}
	s.top = newNode
	fmt.Println("Push = ", *s)
}

func (s *stack) Pop() {
	if s.top == nil {
		fmt.Println("stack is empty")
		return
	}

	s.top = s.top.next 
	

}
