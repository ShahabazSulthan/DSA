package main

import "math"

func IsBST(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	if node.data <= min || node.data >= max {
		return false
	}

	return IsBST(node.left, min, node.data) && IsBST(node.right, node.data, max)
}

func (n *Node) isBST() bool {
	return IsBST(n, math.MinInt64, math.MaxInt64)
}
