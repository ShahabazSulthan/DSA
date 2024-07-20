package main

func (list *DoublyLinkedList) insertAtBack(n *Node) {
	if list.head == nil {
		list.insertAtFront(n)
	} else {
		n.prev = list.tail
		list.tail.next = n
		list.tail = n
		list.length++
	}
}
