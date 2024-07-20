package main

import "fmt"

func (list *DoublyLinkedList) DeleteAtBack() {
	if list.head == nil {
		fmt.Println("linked list is empty")
	} else if list.length == 1 {
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}
	list.length--
	fmt.Println("Delete at Tail Node")
}
