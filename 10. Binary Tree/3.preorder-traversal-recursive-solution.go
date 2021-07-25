package main

import (
	"./util"
	"fmt"
)

func preOrderTraversal(root *util.Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Data, " ")
	preOrderTraversal(root.Left)
	preOrderTraversal(root.Right)
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	fmt.Println("\nThe PreOrder Traversal")
	preOrderTraversal(t.Root)
	fmt.Println()
}
