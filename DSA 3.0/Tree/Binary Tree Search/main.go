package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (n *Node) insert(Value int) {
	if Value <= n.data {
		if n.left == nil {
			n.left = &Node{data: Value}
		} else {
			n.left.insert(Value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: Value}
		} else {
			n.right.insert(Value)
		}
	}
}

func (n *Node) InOrderTraversal() {
	if n.left != nil {
		n.left.InOrderTraversal()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.InOrderTraversal()
	}
}

func (n *Node) PreOrderTraversal() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.PreOrderTraversal()
	}
	if n.right != nil {
		n.right.PreOrderTraversal()
	}
}

func (n *Node) PostOrderTraversal() {
	if n.left != nil {
		n.left.PostOrderTraversal()
	}

	if n.right != nil {
		n.right.PostOrderTraversal()
	}

	fmt.Println(n.data)
}

func (n *Node) LevelOrderTraversal() {
	if n == nil {
		return
	}

	queue := list.New()
	queue.PushBack(n)

	for queue.Len() > 0 {

		element := queue.Front()
		currentNode := element.Value.(*Node)
		fmt.Println(currentNode.data, " ")

		queue.Remove(element)

		if currentNode.left != nil {
			queue.PushBack(currentNode.left)
		}

		if currentNode.right != nil {
			queue.PushBack(currentNode.right)
		}
	}
}

func main() {
	root := &Node{data: 10}
	root.insert(3)
	root.insert(5)
	root.insert(15)
	root.insert(8)
	root.insert(17)
	root.insert(4)
	root.insert(19)

	fmt.Println("In-Order-Traversal")
	root.InOrderTraversal()

	fmt.Println("Pre-Order-Traversal")
	root.PreOrderTraversal()

	fmt.Println("Post-Order-Traversal")
	root.PostOrderTraversal()

	fmt.Println("LevelOrder Traversal:")
	root.LevelOrderTraversal()

	fmt.Println("\nDeleting 15 and 5 from the tree:")
	root = root.Delete(15)
	root = root.Delete(5)

	fmt.Println("In-Order-Traversal after deletion:")
	root.InOrderTraversal()

	fmt.Println("Height of the tree:")
	fmt.Println(root.Height())

	fmt.Println("Is Tree a BST?")
	fmt.Println(root.isBST())

	fmt.Println("Degree of the root node:")
	fmt.Println(root.Degree())

	fmt.Println("Degrees of all nodes:")
	root.PrintAllDegree()

	fmt.Println("Level Order")
	root.LevelOrder()

	var values []int
	root.InOrder(&values)

	secondLarget := values[len(values)-2]
	secondSmallest := values[1]

	fmt.Println("\nSecond Largest = ", secondLarget)
	fmt.Println("Second Smallest = ", secondSmallest)

	fmt.Println("Depth = ",root.Depth())

	// Example for finding height and depth of specific nodes:
	target := 8
	node := findNode(root, target)
	if node != nil {
		fmt.Printf("Height of node with value %d: %d\n", target, node.Height())
		fmt.Printf("Depth of node with value %d: %d\n", target, node.Depth())
	} else {
		fmt.Printf("Node with value %d not found.\n", target)
	}
}
