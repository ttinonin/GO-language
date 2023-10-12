package main

import (
	"fmt"
)

type Node struct {
	value int
	next *Node
} 

type List struct {
	size int
	head *Node
}

func init_list() *List  {
	l := &List{head: nil, size: 0}
	return l
}

func remove(l *List, value int) {
	curr := l.head
	var trash *Node

	if curr != nil && l.head.value == value {
		trash = curr
		l.head = trash.next
	} else {
		for curr != nil && curr.next != nil && curr.next.value != value {
			curr = curr.next
		}

		if curr != nil && curr.next != nil {
			trash = curr.next
			curr.next = trash.next
		}
	}

	if trash != nil {
		l.size--
	}
}

func insert(l *List, value int) {
	node := &Node{value: value, next: nil}

	if l.head == nil {
		l.head = node
	} else {
		var prev *Node

		curr := l.head
		
		for curr != nil && curr.value <= value {
			prev = curr
			curr = curr.next
		}

		if prev == nil {
			node.next = l.head
			l.head = node
		} else {
			if curr == nil {
				prev.next = node
			} else {
				node.next = curr
				prev.next = node
			}
		}
	}

	l.size++
}

func show(l *List) {
	node := l.head

	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.next
	}
	fmt.Printf("\n")
}

func main() {
	list := init_list()

	insert(list, 12)
	insert(list, 24)
	insert(list, 3)
	insert(list, 7)
	insert(list, 32)
	insert(list, 5)

	remove(list, 12)

	show(list)
}
