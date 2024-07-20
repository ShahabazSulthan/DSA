package main

import "fmt"

func (list *DoublyLinkedList) DeleteAtFront() {
	if list.length == 1 {
		list.head = nil
		list.tail = nil
	} else {
		list.head = list.head.next
		list.head.prev = nil
	}
	list.length--
	fmt.Println("delete at head node")
}
