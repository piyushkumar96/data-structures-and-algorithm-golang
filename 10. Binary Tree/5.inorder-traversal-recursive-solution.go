package main

import (
	"./util"
	"fmt"
)

func inOrderTraversal(root *util.Node) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)
	fmt.Print(root.Data, " ")
	inOrderTraversal(root.Right)
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	fmt.Println("\nThe InOrder Traversal")
	inOrderTraversal(t.Root)
	fmt.Println()
}
