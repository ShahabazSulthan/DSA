package main

import "fmt"

func (list *DoublyLinkedList) deleteAtPos(pos int) {
	if pos > list.length {
		fmt.Println("Size exceeding")
	} else {
		if pos == 0 {
			list.DeleteAtFront()
		} else if pos == -1 {
			list.DeleteAtBack()
		}
		current := list.head
		index := 0
		for index < pos {
			index++
			current = current.next
		}
		if current == list.tail {
			list.DeleteAtBack()
		} else {
			current.next.prev = current.prev
			current.prev.next = current.next
			list.length--
		}
	}
}
