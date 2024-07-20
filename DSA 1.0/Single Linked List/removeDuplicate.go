package main



// removeDuplicate removes all duplicates from the linked list.
func (list *LinkedList) removeDuplicate() {
	current := list.head
	var prev *Node
	seen := make(map[int]bool)

	for current != nil {
		if seen[current.data] {
			prev.next = current.next
		} else {
			seen[current.data] = true
			prev = current
		}
		current = current.next
	}
}
