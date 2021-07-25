package main

import (
	"./util"
	"fmt"
)

func preOrderTraversalIterative(root *util.Node) {
	if root == nil {
		return
	}

	st := make([]*util.Node, 0)
	st = append(st, root)
	for len(st) != 0 {
		node := st[len(st)-1]
		fmt.Print(node.Data, " ")
		st = st[:len(st)-1]
		if node.Right != nil {
			st = append(st, node.Right)
		}
		if node.Left != nil {
			st = append(st, node.Left)
		}
	}
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	fmt.Println("\nThe PreOrder Traversal")
	preOrderTraversalIterative(t.Root)
	fmt.Println()
}
