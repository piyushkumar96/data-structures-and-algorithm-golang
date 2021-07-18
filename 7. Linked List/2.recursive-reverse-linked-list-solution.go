package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) insertLast(item interface{}) {
	node := &Node{
		data: item,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
}

func (l *List) printList() {
	if l.head == nil {
		fmt.Println("The list is empty")
		return
	}
	curr := l.head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}

func recReverseList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	rest := recReverseList(head.next)
	head.next.next = head
	head.next = nil

	return rest
}

func main() {
	var l List
	l.insertLast(10)
	l.insertLast(20)
	l.insertLast(30)
	l.insertLast(40)
	l.insertLast(50)
	l.printList()

	l.head = recReverseList(l.head)
	l.printList()

}
