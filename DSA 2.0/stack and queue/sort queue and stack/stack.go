package main

func reverseStack(stack []int) []int {
	var queue []int

	for len(stack) > 0 {
		element := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		queue = append(queue, element)
	}

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		stack = append(stack, element)
	}

	return stack
}