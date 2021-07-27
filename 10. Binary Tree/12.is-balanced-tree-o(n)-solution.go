package main

import (
	"./util"
	"fmt"
	"math"
)

type HBPair struct {
	height   int
	balanced bool
}

func isBalanced_ON(root *util.Node) HBPair {
	var hbp HBPair
	if root == nil {
		hbp.height = 0
		hbp.balanced = true
		return hbp
	}

	lpa := isBalanced_ON(root.Left)
	rpa := isBalanced_ON(root.Right)

	hbp.height = int(math.Max(float64(lpa.height), float64(rpa.height))) + 1

	if int(math.Abs(float64(lpa.height)-float64(rpa.height))) <= 1 && lpa.balanced && rpa.balanced {
		hbp.balanced = true
		return hbp
	}

	hbp.balanced = false
	return hbp
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	pa := isBalanced_ON(t.Root)

	if pa.balanced {
		fmt.Println("\nThe Tree is balanced")
	} else {
		fmt.Println("\nThe Tree is not balanced")
	}
}
