package main

import (
	"./util"
	"fmt"
)

func inOrderTraversalIterative(root *util.Node) {
	if root == nil {
		return
	}

	st := make([]*util.Node, 0)
	node := root
	for node != nil || len(st) != 0 {
		for node != nil {
			st = append(st, node)
			node = node.Left
		}

		node = st[len(st)-1]
		fmt.Print(node.Data, " ")
		st = st[:len(st)-1]

		node = node.Right
	}
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	fmt.Println("\nThe InOrder Traversal")
	inOrderTraversalIterative(t.Root)
	fmt.Println()
}
