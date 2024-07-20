package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Push(val int) {
	s.elements = append(s.elements, val)
}

func (s *Stack) Pop() int {
	if s.isEmpty() {
		fmt.Println("stack is empty")
		return 0
	}

	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func (s *Stack) Peek() int {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return 0
	}

	top := s.elements[len(s.elements)-1]
	return top
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func main() {
	s := Stack{}

	s.Push(90)
	s.Push(60)
	s.Push(30)

	fmt.Println("Stack = ",s.elements)

	fmt.Println("Size of stack = ",s.Size())

	fmt.Println("poped Top of stack ",s.Pop())

	fmt.Println("Stack = ",s.elements)

	fmt.Println("Peek of stack",s.Peek())

	fmt.Println("Is Stack Empty = ",s.isEmpty())

}
