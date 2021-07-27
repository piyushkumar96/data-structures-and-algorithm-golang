package main

import (
	"./util"
	"fmt"
	"math"
)

func heightOfTree(root *util.Node) int {
	if root == nil {
		return 0
	}

	lh := heightOfTree(root.Left)
	rh := heightOfTree(root.Right)

	return int(math.Max(float64(lh), float64(rh))) + 1
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	h := heightOfTree(t.Root)
	fmt.Println("\nThe Height of Tree is ", h)
}
