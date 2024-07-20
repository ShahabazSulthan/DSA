package main

import "fmt"

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
	fmt.Println("Push = ",*s)
}

func (s *Stack) Pop() {
	if len(*s) == 0 {
		fmt.Println("Stack is empty")
		return
	} else {
		top := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]

		fmt.Println("Top = ",top," POP = ",*s)
	}
}

func (s *Stack) Peek() {
	if len(*s) == 0 {
		fmt.Println("Stack is empty")
		return
	} else {
		peek := (*s)[len(*s)-1]
		fmt.Println("PEEK = ",peek)
	}
}

func (s *Stack) isEmpty() {
	fmt.Println("IsEmpty = ",len(*s) == 0 )
}

func main() {
	s := Stack{}

	s.Push(22)
	s.Push(44)
	s.Push(55)

	s.Pop()
	s.Peek()
	s.isEmpty()
}
