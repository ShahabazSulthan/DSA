package main

import "fmt"

func reveseQueue(queue []int) []int {
	var stack []int

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		stack = append(stack, element)
	}

	for len(stack) > 0 {
		element := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		queue = append(queue, element)
	}
	return queue
}

func main() {
	queue := []int{66, 77, 88, 99}

	fmt.Println("Revised Queue = ",reveseQueue(queue))

	stack := []int{999,555,333,222,111}

	fmt.Println("Revised Stack = ",reverseStack(stack))


}
