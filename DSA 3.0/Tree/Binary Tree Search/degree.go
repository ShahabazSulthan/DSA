package main

import "fmt"

func (n *Node) Degree() int {
	if n == nil {
		return 0
	}

	degree := 0

	if n.left != nil {
		degree++
	}

	if n.right != nil {
		degree++
	}

	return degree
}

func (n *Node) PrintAllDegree() {
	if n == nil {
		return
	}

	fmt.Printf("Node = %d  Degree = %d\n", n.data, n.Degree())

	if n.left != nil {
		n.left.PrintAllDegree()
	}

	if n.right != nil {
		n.right.PrintAllDegree()
	}
}
