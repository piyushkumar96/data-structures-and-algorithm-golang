package main

import (
	"./util"
	"fmt"
)

func levelOrderBuildTree() *util.Node {
	var q util.Queue
	var item int
	_, _ = fmt.Scan(&item)

	root := &util.Node{
		Data:  item,
		Left:  nil,
		Right: nil,
	}

	q.Push(root)

	for !q.IsEmpty() {
		var c1, c2 int
		_, _ = fmt.Scan(&c1, &c2)

		node, _ := q.Pop()
		if c1 != -1 {
			node.Left = &util.Node{
				Data:  c1,
				Left:  nil,
				Right: nil,
			}
			q.Push(node.Left)
		}

		if c2 != -1 {
			node.Right = &util.Node{
				Data:  c2,
				Left:  nil,
				Right: nil,
			}
			q.Push(node.Right)
		}
	}

	return root
}

func main() {
	var t util.Tree
	t.Root = levelOrderBuildTree()
	t.PrintTree()
}
