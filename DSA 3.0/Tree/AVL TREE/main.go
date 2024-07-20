package main

import "fmt"

type Node struct {
	key    int
	height int
	right  *Node
	left   *Node
}

func height(n *Node) int {
	if n == nil {
		return 0
	}

	return n.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Create a new node
func newNode(key int) *Node {
	return &Node{key: key, height: 1}
}

// Get the balance factor of node
func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func main() {
	root := insert(nil, 10)
	root = insert(root, 20)
	root = insert(root, 30)
	root = insert(root, 40)
	root = insert(root, 50)
	root = insert(root, 25)

	fmt.Print("In-order traversal of the constructed AVL tree is: ")
	inOrder(root)
}
