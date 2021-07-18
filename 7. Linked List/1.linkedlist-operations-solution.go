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

func (l *List) insertFront(item interface{}) {
	node := &Node{
		data: item,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *List) deleteLast() {
	if l.head == nil {
		fmt.Println("List is underFlow")
		return
	}
	if l.head == l.tail {
		l.tail = nil
		l.head = nil
		return
	}
	curr := l.head
	for curr.next != l.tail {
		curr = curr.next
	}

	curr.next = nil
	l.tail = curr
}

func (l *List) deleteFront() {
	if l.head == nil {
		fmt.Println("List is UnderFlow")
		return
	}
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	l.head = l.head.next
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

func main() {
	var l List
	l.deleteLast()
	l.deleteFront()
	l.insertLast(10)
	l.printList()
	l.deleteLast()
	l.printList()
	l.insertLast(20)
	l.printList()
	l.deleteFront()
	l.printList()
	l.insertLast(30)
	l.insertLast(40)
	l.insertFront(5)
	l.printList()
	l.deleteLast()
	l.printList()
	l.insertLast(55)
	l.insertFront(2)
	l.printList()
	l.deleteFront()
	l.printList()

}
