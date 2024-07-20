package main

import "fmt"

type Node struct {
	data int   // Data stored in the node
	next *Node // Pointer to the next node
	prev *Node // Pointer to the previous node
}

type DoublyLinkedList struct {
	head   *Node // Pointer to the first node
	tail   *Node // Pointer to the last node
	length int   // Length of the linked list
}

func main() {

	List := &DoublyLinkedList{}

	node1 := &Node{data: 11}
	node2 := &Node{data: 22}
	node3 := &Node{data: 44}
	node4 := &Node{data: 33}

	List.insertAtBack(node3)
	List.insertAtFront(node2)
	List.insertAtFront(node1)

	List.DisplayList()

	fmt.Println("Length of list = ", List.FindLength())

	fmt.Println(List.FindElement(11))

	List.InsertatPos(node4, 2)

	List.DisplayList()

	List.DeleteAtFront()

	List.DisplayList()

	List.DeleteAtBack()

	List.DisplayList()

	List.insertAtFront(node3)

	List.DisplayList()

	List.deleteAtPos(1)

	List.DisplayList()

	List.deleteAllList()

	List.DisplayList()

}
