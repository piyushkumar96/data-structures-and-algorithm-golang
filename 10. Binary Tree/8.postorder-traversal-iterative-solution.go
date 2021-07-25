package main

import (
	"./util"
	"fmt"
)

func postOrderTraversalIterative(root *util.Node) {
	if root == nil {
		return
	}
	st := make([]*util.Node, 0)

	node := root
	for ok := true; ok; ok = len(st) > 0 {
		for node != nil {
			if node.Right != nil {
				st = append(st, node.Right)
			}
			st = append(st, node)
			node = node.Left
		}
		node = st[len(st)-1]
		st = st[:len(st)-1]

		if node.Right != nil && len(st) > 0 && st[len(st)-1] == node.Right {
			st = st[:len(st)-1]
			st = append(st, node)
			node = node.Right
		} else {
			fmt.Print(node.Data, " ")
			node = nil
		}
	}
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	fmt.Println("\nThe PostOrder Traversal")
	postOrderTraversalIterative(t.Root)
	fmt.Println()
}
