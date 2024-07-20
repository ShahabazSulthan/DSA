package main

import "fmt"

type Nodes struct {
	data rune
	next *Nodes
}

type Stacks struct {
	top *Nodes
}

func (s *Stacks) Push(val rune) {
	node := &Nodes{data: val, next: s.top}
	s.top = node
}

func (s *Stacks) pop() rune {
	if s.top == nil {
		fmt.Println("stack is empty")
		return 0
	}
	top := s.top.data
	s.top = s.top.next
	return rune(top)
}

// isEmpty checks if the stack is empty.
func (s *Stacks) isEmpty() bool {
	return s.top == nil
}

func reverseString(input string) string {
	stack := &Stacks{}

	for _, c := range input {
		stack.Push(c)
	}

	var reverse []rune
	for !stack.isEmpty() {
		reverse = append(reverse, stack.pop())
	}
	return string(reverse)
}

