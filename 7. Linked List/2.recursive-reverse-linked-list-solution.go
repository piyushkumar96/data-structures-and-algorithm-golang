package main

import (
	"./util"
)

func recReverseList(head *util.Node) *util.Node {
	if head == nil || head.Next == nil {
		return head
	}

	rest := recReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return rest
}

func main() {
	l := util.CreateList()
	l.InsertLast(10)
	l.InsertLast(20)
	l.InsertLast(30)
	l.InsertLast(40)
	l.InsertLast(50)
	l.PrintList()

	l.Head = recReverseList(l.Head)
	l.PrintList()

}
