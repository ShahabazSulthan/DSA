package main

import "fmt"

func (list *LinkedList) insertBeforeNode(beforeValue, data int) {
	if list.head.data == beforeValue {
		newNode := &Node{data: data, next: list.head}
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil {
		if current.data == beforeValue {
			newNode := &Node{data: data, next: current.next}
			current.next = newNode
			return
		}
		current = current.next
	}
	fmt.Printf("Cannot insert node, after value %d doesn't exist", beforeValue)
	fmt.Println("")
}
