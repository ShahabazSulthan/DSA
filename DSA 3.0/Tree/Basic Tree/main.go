package main

import "fmt"

type TreeNode struct {
	data     int
	children []*TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		data:     value,
		children: []*TreeNode{},
	}
}

func (t *TreeNode) insert(value int) {
	child := NewTreeNode(value)
	t.children = append(t.children, child)
}

func Spaces(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += "--"
	}
	return result
}

func printTree(node *TreeNode, level int) {
	if node == nil {
		return
	}

	fmt.Printf("%s%d|\n", Spaces(level), node.data)

	for _, child := range node.children {
		printTree(child, level+1)
	}
}

func (t *TreeNode) PrintTree() {
	printTree(t, 0)
}

func main() {
	root := NewTreeNode(1)

	root.insert(2)
	root.insert(3)
	root.children[0].insert(4)
	root.children[0].insert(5)
	root.children[1].insert(6)

	fmt.Println(".....Tree.....")
	root.PrintTree()
}
