package main

import (
	"./util"
	"fmt"
)

func replaceWithDescendantsSum(root *util.Node) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Data.(int)
	}

	lsum := replaceWithDescendantsSum(root.Left)
	rsum := replaceWithDescendantsSum(root.Right)

	temp := root.Data
	root.Data = lsum + rsum

	return temp.(int) + lsum + rsum
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	replaceWithDescendantsSum(t.Root)
	fmt.Println("\nThe Sum Tree is ")
	t.PrintTree()
}
