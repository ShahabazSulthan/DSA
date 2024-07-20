package main

import "fmt"

func (list *LinkedList) insertAfterNode(aftervalue, data int) {
	newNode := &Node{data: data, next: nil}
	current := list.head

	for current != nil {
		if current.data == aftervalue {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}

	fmt.Printf("Cannot insert node, after value %d doesn't exist", aftervalue)
	fmt.Println("")
}
