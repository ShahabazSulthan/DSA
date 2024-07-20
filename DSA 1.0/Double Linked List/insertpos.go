package main

import "fmt"

func (list *DoublyLinkedList) InsertatPos(n *Node, pos int) {
	if list.length < pos {
		fmt.Println("size exceed")
	} else if pos == 0 {
		list.insertAtFront(n)
	} else if pos == -1 {
		list.insertAtBack(n)
	} else {
		current := list.head
		index := 0
		for index < pos-1 {
			current = current.next
			index++
		}
		current.next.prev = n
		n.next = current.next
		n.prev = current
		current.next = n
		list.length++
	}
}
