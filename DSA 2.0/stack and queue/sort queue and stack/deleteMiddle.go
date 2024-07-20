package main

import "fmt"

type Stack struct {
	element []int
}

func (s *Stack) Push(val int) {
	s.element = append(s.element, val)
}

func (s *Stack) Pop() int {
	if len(s.element) == 0 {
		fmt.Println("Stack is empty")
		return 0
	}

	top := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]
	return top
}

func DeleteMiddle(s *Stack,currIndex,size int) {
	if currIndex == size/2 {
		s.Pop()
		return
	}
	value := s.Pop()
	DeleteMiddle(s,currIndex+1,size)
	s.Push(value)
}

func deleteSpecific(s *Stack,value int) {
    if len(s.element) <= 0 {
		fmt.Println("Stack is empty")
		return
	}

	top := s.Pop()

	if top == value {
		return
	}

	deleteSpecific(s,value)

	s.Push(top)
}

func main() {
	s := &Stack{}

	arr := []int{11,22,33,44,55,66,77}

	for _,v := range arr {
		s.Push(v)
	}

	
	fmt.Println("Stack before deleting middle element:", s.element)

	// Deleting the middle element
	DeleteMiddle(s, 0, len(s.element))

	fmt.Println("Stack after deleting middle element:", s.element)

	deleteSpecific(s,33)

	fmt.Println("stack is after deleting specific pos:",s.element)
}
