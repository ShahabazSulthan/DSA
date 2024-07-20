package main

import "fmt"

func (n *Node) LevelOrder() {
	if n == nil {
		return
	}

	queue := []*Node{n}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		fmt.Printf("%d  ", currentNode.data)

		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}

		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}
}
