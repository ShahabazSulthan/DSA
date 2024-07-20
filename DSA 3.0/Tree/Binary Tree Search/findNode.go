package main

func findNode(n *Node, value int) *Node {
	if n == nil || n.data == value {
		return n
	}

	if value < n.data {
		return findNode(n.left, value)
	} else {
		return findNode(n.right, value)
	}
}
