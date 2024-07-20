package main

func (list *LinkedList) arrayTOLinkedlist(arr []int) *LinkedList {
	if len(arr) == 0 {
		return &LinkedList{}
	}

	head := &Node{data: arr[0], next: nil}
	current := head

	for _, value := range arr[1:] {
		newNode := &Node{data: value, next: nil}
		current.next = newNode
		current = newNode
	}
	
	list.head = head
	return list
}
