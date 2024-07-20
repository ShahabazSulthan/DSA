package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	Mylist := LinkedList{}

	Mylist.insertAtFront(20)
	Mylist.insertAtFront(30)
	Mylist.insertAtFront(40)
	Mylist.insertAtFront(50)

	Mylist.print()

	fmt.Println("Count Of The LinkedList = ", Mylist.countNode())

	Mylist.insertAtback(10)

	Mylist.insertAtFront(60)

	Mylist.print()

	Mylist.insertAfterNode(50, 55)

	Mylist.print()

	Mylist.insertAfterNode(70, 5)

	Mylist.insertBeforeNode(40, 35)

	Mylist.print()

	//Mylist.insertBeforeNode(80, 86)

	Mylist.deleteFrontNode()

	Mylist.print()

	Mylist.deleteLastNode()

	Mylist.print()

	Mylist.deleteAfterNode(40)

	Mylist.print()

	Mylist.deleteBeforeNode(40)

	Mylist.print()

	fmt.Println("head = ", Mylist.head.data)

	fmt.Println("index 3 node = ", Mylist.findNodeAt(3))

	fmt.Println("Count Of The LinkedList = ", Mylist.countNode())

	Mylist.reverse()

	Mylist.print()

	Mylist.sort()

	Mylist.print()

	Mylist.reverse()

	Mylist.print()

	Mylist.sort()

	Mylist.print()

	var arr = []int{5, 4, 4, 3, 1, 2, 1}

	Mylist.arrayTOLinkedlist(arr)

	Mylist.print()

	Mylist.removeDuplicate()

	Mylist.print()

	Mylist.sort()

	Mylist.print()


}
