package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(val int) {
	node := &Node{data: val, next: s.top}
	s.top = node
	s.Print()
}

func (s *Stack) Pop() int {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return 0
	}
	top := s.top
	s.top = s.top.next
	s.Print()
	return top.data
}

func (s *Stack) Peek() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	peek := s.top.data
	fmt.Println("Top = ", peek)
}

func (s *Stack) Size() int {
	size := 0
	current := s.top
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (s *Stack) Print() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	current := s.top
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (s *Stack) reverse() {
	var next, prev *Node
	current := s.top
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	s.top = prev
	fmt.Print("Reversed Stack = ")
	s.Print()
}

func (s *Stack) deleteMiddleElement() {
	if s.isEmpty() {
		fmt.Println("stack is empty")
		return
	}
	size := s.Size()
	mid := size / 2
	s.recursivelyFindMiddle(s.top, mid)
	s.Print()
}

func (s *Stack) recursivelyFindMiddle(current *Node, index int) *Node {
	if index == 0 {
		return current.next
	}
	current.next = s.recursivelyFindMiddle(current.next, index-1)
	return current
}

func SortStack(s *Stack) {
	if !s.isEmpty() {
		top := s.Pop()
		SortStack(s)
		insertInSortOrder(s, top)
	}
}

func insertInSortOrder(s *Stack, item int) {
	if s.isEmpty() || s.top.data > item {
		s.Push(item)
	} else {
		top := s.Pop()
		insertInSortOrder(s, item)
		s.Push(top)

	}
}

func (s *Stack) deletePos(pos int) {
	if pos < 0 || pos >= s.Size() {
		fmt.Println("Position is out of stack")
		return
	}

	temp := Stack{}

	// Pop elements until the desired position
	for i := 0; i < pos; i++ {
		temp.Push(s.Pop())
	}

	// Pop the target element (deleting it)
	s.Pop()

	// Push the elements back from the temporary stack
	for !temp.isEmpty() {
		s.Push(temp.Pop())
	}
}
func main() {
	s := Stack{}

	s.Push(11)
	s.Push(22)
	s.Push(44)
    fmt.Println("Stack before deletion:")
	s.Print()

	fmt.Println("Deleting element at position 2")
	s.deletePos(2)

	fmt.Println("Stack after deletion:")
	s.Print()

	stack := &Stack{}
	stack.Push(3)
	stack.Push(1)
	stack.Push(4)
	stack.Push(2)

	SortStack(stack)

	stack.Print()

	s.Size()

	s.Peek()

	s.Pop()

	fmt.Println("Size = ", s.Size())

	s.Push(55)

	s.reverse()

	fmt.Println("Delete Middle Element:")
	s.deleteMiddleElement()

	input := "Hello, World!"
	fmt.Println("Original String:", input)
	reversed := reverseString(input)
	fmt.Println("Reversed String:", reversed)

}
