package main

// Left rotate
func leftRotate(x *Node) *Node {
	y := x.right
	T2 := y.left
	y.left = x
	x.right = T2

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}
