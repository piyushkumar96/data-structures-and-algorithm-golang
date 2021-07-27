package main

import (
	"./util"
	"fmt"
	"math"
)

func isBalanced_ONN(root *util.Node) bool {
	if root == nil {
		return true
	}

	lh := util.HeightOfTree(root.Left)
	rh := util.HeightOfTree(root.Right)

	if int(math.Abs(float64(lh)-float64(rh))) <= 1 && isBalanced_ONN(root.Left) && isBalanced_ONN(root.Right) {
		return true
	}

	return false
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	isBal := isBalanced_ONN(t.Root)

	if isBal {
		fmt.Println("\nThe Tree is balanced")
	} else {
		fmt.Println("\nThe Tree is not balanced")
	}
}
