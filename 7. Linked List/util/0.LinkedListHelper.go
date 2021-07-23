package util

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func CreateList() *List {
	return &List{
		Head: nil,
		Tail: nil,
	}
}

func (l *List) InsertLast(item interface{}) {
	node := &Node{
		Data: item,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *List) InsertFront(item interface{}) {
	node := &Node{
		Data: item,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

func (l *List) DeleteLast() {
	if l.Head == nil {
		fmt.Println("List is underFlow")
		return
	}
	if l.Head == l.Tail {
		l.Tail = nil
		l.Head = nil
		return
	}
	curr := l.Head
	for curr.Next != l.Tail {
		curr = curr.Next
	}

	curr.Next = nil
	l.Tail = curr
}

func (l *List) DeleteFront() {
	if l.Head == nil {
		fmt.Println("List is UnderFlow")
		return
	}
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		return
	}

	l.Head = l.Head.Next
}

func (l *List) PrintList() {
	if l.Head == nil {
		fmt.Println("The list is empty")
		return
	}
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.Data, " ")
		curr = curr.Next
	}
	fmt.Println()
}
