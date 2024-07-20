package main

import "fmt"

func (list *DoublyLinkedList) DisplayList() {
	if list.head == nil {
		fmt.Println("Linked list is empty")
	} else {
		current := list.head
		for current != nil {
			fmt.Printf("%d -> ", current.data)
			current = current.next
		}
		fmt.Println("END")
	}
}
