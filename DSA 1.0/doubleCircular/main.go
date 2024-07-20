package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *LinkedList) insertFront(n *Node) {
	if list.head == nil {
		list.head = n
		list.tail = list.head
		// Make the list circular
		list.head.next = list.tail
		list.tail.prev = list.head
		list.tail.next = list.head
		list.head.prev = list.tail
	} else {
		n.next = list.head
		n.prev = list.tail
		list.tail.next = n
		list.head.prev = n
		list.head = n
	}
	list.length++
	list.Display()
}

func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	} else {
		temp := list.head
		for temp != nil {
			fmt.Printf("%d -> ", temp.data)
			temp = temp.next
			if temp == list.tail.next {
				break
			}
		}
		fmt.Println("END")
	}
}

func (list *LinkedList) insertEnd(n *Node) {
	if list.head == nil {
		list.insertFront(n)
	} else {
		list.tail.next = n
		n.prev = list.tail
		list.tail = n
		list.tail.next = list.head
		list.head.prev = list.tail
		list.length++
	}
	list.Display()
}

func (list *LinkedList) inserPos(n *Node, pos int) {
	if pos == 0 {
		list.insertFront(n)
	} else if pos == -1 {
		list.insertEnd(n)
	} else {
		temp := list.head
		index := 0
		for index < pos-1 {
			index++
			temp = temp.next
		}
		n.next = temp.next
		temp.next = n
		n.prev = temp
		list.length++
		list.Display()
	}

}

func (list *LinkedList) deleteFront() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	} else {
		list.head = list.head.next
		list.head.prev = list.tail
		list.tail.next = list.head
		list.length--
	}
	list.Display()
}

func (list *LinkedList) deleteEnd() {
	if list.head == nil {
		list.deleteFront()
	} else {
		list.tail = list.tail.prev
		list.head.prev = list.tail
		list.length--
	}
	list.Display()
}

func (list *LinkedList) deletePos(pos int) {
	if pos == 0 {
		list.deleteFront()
		return
	} else if pos == -1 {
		list.deleteEnd()
		return
	} else {
		temp := list.head
		index := 0
		for index < pos {
			index++
			temp = temp.next
		}
		temp.prev.next = temp.next
		temp.next.prev = temp.prev
		list.length--
	}
	list.Display()
}

// // Function to convert an array to a linked list
// func arrayToList(arr []int) {
// 	list := &LinkedList{}

// 	if len(arr) == 0 {
// 		fmt.Println("List is empty")
// 		return
// 	}

// 	for _, val := range arr {
// 		list.insertEnd(&Node{data: val})
// 	}

// 	list.Display()
// }

func (list *LinkedList) revese() {
	var next, prev *Node
	current := list.head
	for {
		next = current.next
		current.next = prev
		prev = current
		current = next
		if current == list.head {
			break
		}
	}
	list.tail = list.head
	list.head = prev

	list.head.prev = list.tail
	list.tail.next = list.head
	list.Display()
}

func main() {
	myList := LinkedList{}

	node1 := &Node{data: 11}
	node2 := &Node{data: 22}
	node3 := &Node{data: 33}
	node7 := &Node{data: 44}
	node4 := &Node{data: 55}
	node5 := &Node{data: 66}
	node6 := &Node{data: 77}

	myList.insertFront(node1)
	myList.insertFront(node2)
	myList.insertFront(node3)

	myList.insertEnd(node4)
	myList.insertEnd(node5)
	myList.insertEnd(node6)

	myList.inserPos(node7, 3)

	myList.deleteFront()

	myList.deleteEnd()

	myList.deletePos(2)

	// var arr = []int{100,200,300,400,500}

	// arrayToList(arr)

	myList.revese()

}
