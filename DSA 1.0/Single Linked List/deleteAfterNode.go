package main

import "fmt"

func (list *LinkedList) deleteAfterNode(aftervalue int) {
	current := list.head

	for current != nil && current.data != aftervalue {
		current = current.next
	}

	if current == nil {
		fmt.Printf("Node with after value %d doesn't exist\n", aftervalue)
		return
	}

	if current.next == nil {
		fmt.Printf("Node with after value %d is the last node\n", aftervalue)
		return
	}

	temp := current.next

	fmt.Printf("Node after value %d has data %d and will be deleted", aftervalue, temp.data)
	fmt.Println()

	current.next = current.next.next
}
