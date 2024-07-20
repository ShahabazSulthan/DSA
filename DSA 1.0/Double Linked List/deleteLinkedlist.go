package main

import "fmt"

func (list *DoublyLinkedList) deleteAllList() {
	if list.head != nil {
		current := list.head
		for current.next != nil {
			current.prev = nil
			current = current.next
		}
		list.head = nil
		list.length = 0
		fmt.Println("Delete the list")
	}
}
