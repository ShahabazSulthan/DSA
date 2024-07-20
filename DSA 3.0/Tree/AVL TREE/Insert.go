package main


// Insert a key into the AVL tree
func insert(node *Node, key int) *Node {
	if node == nil {
		return newNode(key)
	}

	if key < node.key {
		node.left = insert(node.left, key)
	} else if key > node.key {
		node.right = insert(node.right, key)
	} else {
		return node
	}

	node.height = 1 + max(height(node.left), height(node.right))

	balance := getBalance(node)

	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}
	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}
	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}
	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}
