package main

func (n *Node) Depth() int {
	if n== nil {
		return 0
	}

	depth := 0
	current := n
	for current != nil {
		depth++
		current = current.left
	}
	return depth - 1
}