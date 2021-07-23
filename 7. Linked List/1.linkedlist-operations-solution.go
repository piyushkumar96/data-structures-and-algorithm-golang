package main

import (
	"./util"
)

func main() {
	l := util.CreateList()
	l.DeleteLast()
	l.DeleteFront()
	l.InsertLast(10)
	l.PrintList()
	l.DeleteLast()
	l.PrintList()
	l.InsertLast(20)
	l.PrintList()
	l.DeleteFront()
	l.PrintList()
	l.InsertLast(30)
	l.InsertLast(40)
	l.InsertFront(5)
	l.PrintList()
	l.DeleteLast()
	l.PrintList()
	l.InsertLast(55)
	l.InsertFront(2)
	l.PrintList()
	l.DeleteFront()
	l.PrintList()

}
