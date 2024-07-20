package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type circularLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *circularLinkedList) Display() {
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

func (list *circularLinkedList) insertFront(n *Node) {
	if list.head == nil {
		list.head = n
		list.tail = list.head
		list.tail.next = list.head
	} else {
		n.next = list.head
		list.head = n
		list.tail.next = list.head
	}
	list.length++
	list.Display()
}

func (list *circularLinkedList) insertBack(n *Node) {
	if list.head == nil {
		list.insertFront(n)
	} else {
		list.tail.next = n
		list.tail = n
		list.tail.next = list.head
		list.length++
		list.Display()
	}
}

func (list *circularLinkedList) insertPos(n *Node, pos int) {
	if pos > list.length {
		fmt.Println("size exceeding")
		return
	} else if pos == 0 {
		list.insertFront(n)
		return
	} else if pos == -1 {
		list.insertBack(n)
		return
	} else {
		temp := list.head
		index := 0
		for index < pos-1 {
			temp = temp.next
			index++
		}
		n.next = temp.next
		temp.next = n
		list.length++
		list.Display()
	}

}

func (list *circularLinkedList) deleteFront() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	} else if list.length == 1 {
		list.head = nil
		list.length = 0
	} else {
		list.head = list.head.next
		list.tail.next = list.head
		list.length--
	}
	list.Display()
}

func (list *circularLinkedList) deleteEnd() {
	if list.head == nil {
		fmt.Println("list is empty")
	} else if list.length == 1 {
		list.deleteFront()
	} else {
		temp := list.head
		var prev *Node = nil
		for temp.next != list.head {
			prev = temp
			temp = temp.next
		}
		list.tail = prev
		list.tail.next = list.head
		list.length--
	}
	list.Display()
}

func (list *circularLinkedList) deletePos(pos int) {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	} else if pos == 0 {
		list.deleteFront()
	} else if pos == -1 {
		list.deleteEnd()
	} else {
		temp := list.head
		index := 0
		var prev *Node = nil
		for index < pos {
			prev = temp
			temp = temp.next
			index++
		}
		prev.next = temp.next
		list.length--
	}
	list.Display()
}

func main() {
	myList := circularLinkedList{}

	node1 := &Node{data: 21}
	node2 := &Node{data: 31}
	node3 := &Node{data: 41}
	node4 := &Node{data: 51}
	node5 := &Node{data: 61}
	node6 := &Node{data: 71}
	node7 := &Node{data: 81}

	myList.insertFront(node1)
	myList.insertFront(node2)
	myList.insertFront(node3)

	myList.insertBack(node4)
	myList.insertBack(node5)
	myList.insertBack(node6)

	myList.insertPos(node7, 5)

	myList.deleteFront()

	myList.deleteEnd()

	myList.deletePos(2)

}
