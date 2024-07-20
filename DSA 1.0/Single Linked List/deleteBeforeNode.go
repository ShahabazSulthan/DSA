package main

import "fmt"

func (list *LinkedList) deleteBeforeNode(BeforeValue int) {
	var previous *Node
	current := list.head

	if current == nil || current.next == nil {
		fmt.Printf("Node with before value %d doesn't exist\n", BeforeValue)
		return
	}

	for current.next != nil {
		if current.next.data == BeforeValue {
			if previous == nil {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", BeforeValue, current.data)
				list.head = current.next
			} else {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", BeforeValue, current.data)
				previous.next = current.next
			}
			return
		}
		previous = current
		current = current.next
	}
	fmt.Printf("Node before value %d not found\n", BeforeValue)
}
