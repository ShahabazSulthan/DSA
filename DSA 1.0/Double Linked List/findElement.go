package main

import "fmt"

func (list *DoublyLinkedList) FindElement(val int) string {
	if list.head == nil {
		return "Empty Linked List"
	}

	if list.head.data == val {
		return fmt.Sprintf("%d Is Found", list.head.data)
	} else if list.tail.data == val {
		return fmt.Sprintf("%d Is Found", list.tail.data)
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
			if current.data == val {
				return fmt.Sprintf("%d Is Found", current.data)
			}
		}
	}
	return fmt.Sprintf("%d is not  found", val)
}
