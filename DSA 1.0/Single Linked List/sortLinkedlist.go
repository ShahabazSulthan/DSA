package main

func getMiddle(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	slow := head
	fast := head

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func SortedMerge(a *Node, b *Node) *Node {
    var result *Node
    var last *Node

    for a != nil && b != nil {
        var temp *Node
        if a.data <= b.data {
            temp = a
            a = a.next
        } else {
            temp = b
            b = b.next
        }

        if result == nil {
            result = temp
            last = result
        } else {
            last.next = temp
            last = last.next
        }
    }

    if a != nil {
        last.next = a
    } else {
        last.next = b
    }

    return result
}


func MergerSort(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	middle := getMiddle(head)
	nextToMiddle := middle.next
	middle.next = nil

	left := MergerSort(head)
	right := MergerSort(nextToMiddle)

	sortedList := SortedMerge(right, left)

	return sortedList
}

func (list *LinkedList) sort() {
	list.head = MergerSort(list.head)
}
