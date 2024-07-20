package main

import "fmt"

func (list *LinkedList) deleteFrontNode() {
	if list.head != nil {
		list.head = list.head.next
		fmt.Println("Head node of linked list has been deleted")
		return
	}
}
