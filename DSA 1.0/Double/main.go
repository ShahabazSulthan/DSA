package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type Linkedlist struct {
	head   *Node
	tail   *Node
	length int
}

func (list *Linkedlist) Length() {
	fmt.Println("Length = ", list.length)
}

func (list *Linkedlist) insertFront(n *Node) {
	if list.head == nil {
		list.head = n
		list.tail = n
		n.next = nil
		n.prev = nil
	} else {
		list.head.prev = n
		n.next = list.head
		list.head = n
	}
	list.length++
}

func (list *Linkedlist) insertBack(n *Node) {
	if list.head == nil {
		list.insertFront(n)
	} else if list.head.next == nil {
		list.head.next = n
		list.tail = n
	} else {
		n.prev = list.tail
		list.tail.next = n
		list.tail = n
		list.length++
	}
}

func (list *Linkedlist) insertPos(n *Node, pos int) {
	if pos == 0 {
		list.insertFront(n)
		return
	} else if pos == -1 {
		list.insertBack(n)
		return
	} else {
		current := list.head
		index := 0
		if index < pos-1 {
			index++
			current = current.next
		}
		current.next.prev = n
		n.next = current.next
		n.prev = current
		current.next = n
	}
	list.length++
}

func (list *Linkedlist) deleteFront() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	} else {
		list.head.next.prev = nil
		list.head = list.head.next
		list.length--
	}
}

func (list *Linkedlist) deleteEnd() {
	if list.length == 1 {
		list.deleteFront()
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
		list.length--
	}
}

func (list *Linkedlist) deletePos(pos int) {
	if pos == 0 {
		list.deleteFront()
	} else if pos == -1 {
		list.deleteEnd()
	} else {
		current := list.head
		index := 0
		for index < pos {
			index++
			current = current.next
		}
		current.next.prev = current.next.next.prev
		current.next = current.next.next
		list.length--
	}
}

func (list *Linkedlist) findVal(val int) string {
	if list.head == nil {
		return "List is empty"
	} else if list.head.data == val {
		return fmt.Sprintf("%d is Found\n", list.head.data)
	} else if list.tail.data == val {
		return fmt.Sprintf("%d is Found\n", list.tail.data)
	} else {
		current := list.head
		for current.next != nil {
			if current.data == val {
				return fmt.Sprintf("%d is Found\n", current.data)
			}
			current = current.next
		}
		return fmt.Sprintf("%d is Not Found\n", val)
	}

}

func (list *Linkedlist) revese() {
	var next, prev *Node
	current := list.head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

func (list *Linkedlist) Display() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	}
	current := list.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println("END")
}

func (list *Linkedlist) Sort() {
	if list.length == 1 {
		return
	} else {
		current := list.head
		for current != nil {
			index := current
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
	list.Display()
}

func main() {
	myList := Linkedlist{}

	node := &Node{data: 77}
	node1 := &Node{data: 66}
	node2 := &Node{data: 55}
	node3 := &Node{data: 11}

	myList.insertFront(node)
	myList.insertFront(node1)
	myList.insertBack(node2)

	myList.insertPos(node3, 2)

	myList.Display()

	myList.revese()

	myList.Display()

	myList.Sort()

	myList.Length()

	ans := myList.findVal(77)
	fmt.Println(ans)

	ans = myList.findVal(99)
	fmt.Println(ans)

	myList.deletePos(1)

	myList.Display()

	myList.deleteFront()

	myList.deleteEnd()

	myList.Display()
}
