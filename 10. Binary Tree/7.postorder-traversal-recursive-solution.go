package main

import (
	"./util"
	"fmt"
)

func postOrderTraversal(root *util.Node) {
	if root == nil {
		return
	}
	postOrderTraversal(root.Left)
	postOrderTraversal(root.Right)
	fmt.Print(root.Data, " ")
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	fmt.Println("\nThe PostOrder Traversal")
	postOrderTraversal(t.Root)
	fmt.Println()
}
