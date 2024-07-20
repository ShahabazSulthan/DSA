package main

import "fmt"

func (list *LinkedList) deleteLastNode() {
	if list.head.next == nil {
		list.head = nil
		return
	}

	current := list.head

	for current.next.next != nil {
		current = current.next
	}

	current.next = nil

	fmt.Println("Last node of linked list has been deleted")
}
