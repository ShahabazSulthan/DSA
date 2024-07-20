package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(val int) {
	newNode := &Node{data: val}
	newNode.next = s.top
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() int {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return 0
	}
	top := s.top.data
	s.top = s.top.next
	s.size--
	return top
}

func (s *Stack) Peek() int {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return 0
	}

	return s.top.data
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Print() {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return
	}

	current := s.top 

	for current != nil {
		fmt.Printf("%d -> ",current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	s := Stack{}

	s.Push(55)
	s.Push(11)
	s.Push(99)
	s.Push(33)
	
	s.Print()

	fmt.Println("Size = ",s.size)

	fmt.Println("poped Top element = ",s.Pop())

	s.Print()
	
	fmt.Println("Peek = ",s.Peek())

	fmt.Println("isEmpty = ",s.isEmpty())

	
}
