package main

import "fmt"

type Member struct {
	ID   int
	Name string
	Age  int
}

type Node struct {
	Member Member
	left   *Node
	right  *Node
}

type Bst struct {
	root *Node
}

func insert(n *Node, member Member) *Node {
	if n == nil {
		return &Node{Member: member}
	}

	if member.ID < n.Member.ID {
		n.left = insert(n.left, member)
	} else if member.ID > n.Member.ID {
		n.right = insert(n.right, member)
	}

	return n
}

func (b *Bst) Insert(member Member) {
	b.root = insert(b.root, member)
}

func InOrderTraversal(n *Node) {
	if n != nil {
		InOrderTraversal(n.left)
		fmt.Println("ID = ", n.Member.ID, " Name = ", n.Member.Name, " Age = ", n.Member.Age)
		InOrderTraversal(n.right)
	}
}

func (b *Bst) InOrder() {
	InOrderTraversal(b.root)
}

func delete(node *Node, id int) *Node {
	if node == nil {
		return nil
	}

	if id < node.Member.ID {
		node.left = delete(node.left, id)
	} else if id > node.Member.ID {
		node.right = delete(node.right, id)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		node.Member = minNode(node.right).Member
		node.right = delete(node.right, node.Member.ID)
	}
	return node
}

func minNode(node *Node) *Node {
	currentNode := node
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode
}

func (bst *Bst) Delete(id int) {
	bst.root = delete(bst.root, id)
}

func Search(node *Node,id int) *Member {
	if node == nil {
		return nil
	}

	if id < node.Member.ID {
		return Search(node.left,id)
	} else if id > node.Member.ID {
		return Search(node.right,id)
	}

	return &node.Member
}

func (b *Bst) Search(id int) *Member {
     return Search(b.root,id)
}

func main() {
	bst := &Bst{}

	bst.Insert(Member{ID: 5, Name: "Shahabaz", Age: 21})
	bst.Insert(Member{ID: 3, Name: "Aston", Age: 22})
	bst.Insert(Member{ID: 6, Name: "Emmanuel", Age: 20})
	bst.Insert(Member{ID: 2, Name: "Sujith", Age: 22})

	fmt.Println("In Order Traversal:")
	bst.InOrder()

	fmt.Println("\nDelete member with ID 2")
	bst.Delete(2)
	fmt.Println("In-order traversal after deletion:")
	bst.InOrder()

	fmt.Println("\nSearch for member with ID 6:")
	member := bst.Search(6)
	if member != nil {
		fmt.Printf("Found: ID: %d, Name: %s, Age: %d\n", member.ID, member.Name, member.Age)
	} else {
		fmt.Println("Member not found")
	}
}
