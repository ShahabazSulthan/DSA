package main

func (list *LinkedList) countNode() int {
	count := 0
	current := list.head
	for current != nil {
		current = current.next
		count++
	}
	return count
}
