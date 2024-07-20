package main

// reverse reverses the linked list and updates the head
func (list *LinkedList) reverse() {
	var next, prev *Node
	current := list.head

	// Traverse the list and reverse the links
	for current != nil {
		next = current.next // Store the next node
		current.next = prev // Reverse the link
		prev = current      // Move prev to current node
		current = next      // Move current to next node
	}

	// Update the head to the new front of the list
	list.head = prev
}