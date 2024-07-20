package main

func (n *Node) Delete(Value int) *Node {
	if n == nil {
		return nil
	}

	if Value < n.data {
		n.left = n.left.Delete(Value)
	} else if Value > n.data {
		n.right = n.right.Delete(Value)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		minRight := n.right.findmin()
		n.data = minRight.data
		n.right = n.right.Delete(minRight.data)
	}
	return n
}

func (n *Node) findmin() *Node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}