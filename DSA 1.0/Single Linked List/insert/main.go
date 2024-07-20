package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *LinkedList) insertBeginning(n *Node) {
	if list.head == nil {
		list.head = n
		list.tail = n
	} else {
		n.next = list.head
		list.head = n
	}
	list.length++
}

func (list *LinkedList) insertEnd(n *Node) {
	if list.head == nil {
		list.insertBeginning(n)
	}
	list.tail.next = n
	list.tail = n
	list.length++
}

func (list *LinkedList) insertMiddle(n *Node, pos int) {
	if list.length < pos {
		fmt.Println("size exceeding")
	} else {
		if pos == 0 {
			list.insertBeginning(n)
		} else if pos == -1 {
			list.insertEnd(n)
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
		}

	}

}

func (list *LinkedList) deleteMiddle() {
	if list.head == nil || list.head.next == nil {
		list.head = nil
		return
	}

	slow := list.head
	fast := list.head
	var prev *Node

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		prev = slow
		slow = slow.next
	}
	

	if prev != nil {
		prev.next = slow.next
	}
	list.length--
}

func (list *LinkedList) deleteBegining() {
	if list.head == nil {
		fmt.Println("Linked list is empty....")
		return
	}

	list.head = list.head.next
	list.length--
}

func (list *LinkedList) deleteEnd() {
	if list.head == nil {
		fmt.Println("Linked List is empty")
		return
	} else if list.head.next == nil {
		list.head = nil
		list.length--
		return
	} else {
		temp := list.head
		for temp.next.next != nil {
			temp = temp.next
		}
		temp.next = nil
	}
	list.length--
}

func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("Linked List is empty")
	} else {
		temp := list.head
		for temp != nil {
			fmt.Printf("%v -> ", temp.data)
			temp = temp.next
		}
		fmt.Println("End")
	}
}

func (list *LinkedList) countList() int {
	return list.length
}

func (list *LinkedList) deleteList() {
	list.length = 0
	list.head = nil
	fmt.Println("Deleted All list")
}

func (list *LinkedList) Sort() {
	if list.head == nil {
		return
	}
	current := list.head

	

	for current != nil {
		index := current.next

		for index != nil {
			if current.data > index.data {
				temp := current.data
				current.data = index.data
				index.data = temp
			}
			index = index.next
		}
		current = current.next
	}
}

func (list *LinkedList) reveseList() {
	var next,prev *Node
	current := list.head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

func main() {

	node1 := &Node{data: 6}
	node2 := &Node{data: 88}
	node3 := &Node{data: 9}

	myList := LinkedList{}

	myList.insertBeginning(node1)
	myList.insertEnd(node3)
	myList.insertMiddle(node2, 1)

	myList.Display()

	myList.reveseList()

	myList.Display()

	myList.Sort()

	myList.Display()

	count := myList.countList()
	fmt.Println("List Of Count = ",count)

	myList.deleteMiddle()

	//myList.deleteBegining()
	
	myList.deleteEnd()



	myList.Display()

	myList.deleteList()

	count1 := myList.countList()
	fmt.Println("List Of Count = ",count1)

	myList.Display()

	


}
