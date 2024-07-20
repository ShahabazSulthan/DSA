package main


func (list *LinkedList) insertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data,next: nil}
		list.head = newNode
		return
	}
	newNode := &Node{data: data,next: list.head}
	list.head = newNode
}