package main

func (list *LinkedList) insertAtback(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
