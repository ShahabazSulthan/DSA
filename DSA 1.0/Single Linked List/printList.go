package main

import "fmt"



func (list *LinkedList) print() {
	var current *Node = list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
