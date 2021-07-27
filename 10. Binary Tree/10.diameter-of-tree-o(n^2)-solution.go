package main

import (
	"./util"
	"fmt"
	"math"
)

func diameterOfTree(root *util.Node) int {
	if root == nil {
		return 0
	}

	lh := util.HeightOfTree(root.Left)
	rh := util.HeightOfTree(root.Right)

	d1 := lh + rh + 1
	d2 := diameterOfTree(root.Left)
	d3 := diameterOfTree(root.Right)

	return int(math.Max(math.Max(float64(d1), float64(d2)), float64(d3)))
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	dia := diameterOfTree(t.Root)
	fmt.Println("\nThe Diameter of Tree is ", dia)
}
