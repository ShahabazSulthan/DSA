package main

import "fmt"

type Node struct {
	data rune
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(val rune) {
	newNode := &Node{data: val}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) Pop() rune {
	if s.top == nil {
		fmt.Println("Stack is empmty")
	}

	top := s.top.data
	s.top = s.top.next
	return top
}

func ReverseString(input string) string {
	stack := Stack{}

	for _, char := range input {
		stack.Push(char)
	}

	var revise string
	for !stack.isEmpty() {
		revise += string(stack.Pop())
	}
	return revise
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func main() {
	input := "Shahabaz Sultan"
	fmt.Println("Revised = ",ReverseString(input))

	fmt.Println("Revised Queue = ",reviseQueue(input))
}
