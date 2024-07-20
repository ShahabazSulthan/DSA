package main

import "fmt"

// In-order traversal of the tree
func inOrder(root *Node) {
	if root != nil {
		inOrder(root.left)
		fmt.Printf("%d ", root.key)
		inOrder(root.right)
	}
}
