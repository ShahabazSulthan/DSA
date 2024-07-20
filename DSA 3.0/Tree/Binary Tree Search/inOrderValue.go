package main

func (n *Node) InOrder(value *[]int) {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.InOrder(value)
	}

	*value = append(*value, n.data)

	if n.right != nil {
		n.right.InOrder(value)
	}
}
