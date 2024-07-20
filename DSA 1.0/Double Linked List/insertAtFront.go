package main

func (list *DoublyLinkedList) insertAtFront(n *Node) {
	if list.head == nil {
		list.head = n
		list.tail = n
		n.prev = nil
		n.next = nil
	} else {
		list.head.prev = n
		n.next = list.head
		list.head = n
	}
	list.length++
}
