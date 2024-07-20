package main

func (n *Node) Height() int {
	if n == nil {
		return -1
	}

	leftHeight := n.left.Height()
	rightHeight := n.right.Height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}
