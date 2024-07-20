package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (n *Node) Insert(value int) {
	queue := []*Node{n}

	for len(queue) > 0 {
		curretNode := queue[0]
		queue = queue[1:]

		if curretNode.left == nil {
			curretNode.left = &Node{data: value}
			break
		} else {
			queue = append(queue, curretNode.left)
		}

		if curretNode.right == nil {
			curretNode.right = &Node{data: value}
			break
		} else {
			queue = append(queue, curretNode.right)
		}
	}
}


func (n *Node) InOrder() {
	if n == nil {
		return
	}
	
	if n.left != nil {
		n.left.InOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.InOrder()
	}
}



func main() {
	root := Node{data: 10}
	root.Insert(20)
	root.Insert(30)
	root.Insert(40)
	root.Insert(50)

	fmt.Println("In Order Traversal :")
	root.InOrder()
}
