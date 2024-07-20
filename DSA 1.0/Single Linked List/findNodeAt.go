package main

func (list * LinkedList) findNodeAt(index int) *Node{
  count := 0
  current := list.head

  for current != nil {
	count++
	current = current.next
  }
  
  if index <= 0 || index > count {
	return nil
  }

  current = list.head
  for count = 1;count < index ; count++ {
	current = current.next
  }
  return current
}