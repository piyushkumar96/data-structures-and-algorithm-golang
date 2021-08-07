package main

import (
	"./util"
	"fmt"
	"sort"
)

func printNodesAtKthLevel(root *util.Node, level int, nodes *[]interface{}) {
	if root == nil {
		return
	}

	if level == 0 {
		*nodes = append(*nodes, root.Data)
		return
	}

	printNodesAtKthLevel(root.Left, level-1, nodes)
	printNodesAtKthLevel(root.Right, level-1, nodes)
}

func printNodesFromKthDistance(root *util.Node, target *util.Node, k int, nodes *[]interface{}) int {
	if root == nil {
		return -1
	}

	if root == target {
		printNodesAtKthLevel(root, k, nodes)
		return 0
	}

	DL := printNodesFromKthDistance(root.Left, target, k, nodes)
	if DL != -1 {
		if DL+1 == k {
			*nodes = append(*nodes, root.Data)
		} else {
			printNodesAtKthLevel(root.Right, k-DL-2, nodes)
		}

		return DL + 1
	}

	DR := printNodesFromKthDistance(root.Right, target, k, nodes)
	if DR != -1 {
		if DR+1 == k {
			*nodes = append(*nodes, root.Data)
		} else {
			printNodesAtKthLevel(root.Left, k-DR-2, nodes)
		}

		return DR + 1
	}

	return -1
}

func main() {
	var t util.Tree

	t.Root = util.NewNode(1)
	t.Root.Left = util.NewNode(2)
	t.Root.Right = util.NewNode(3)
	t.Root.Left.Left = util.NewNode(4)
	t.Root.Left.Right = util.NewNode(5)
	t.Root.Left.Right.Left = util.NewNode(7)
	t.Root.Left.Right.Right = util.NewNode(8)
	t.Root.Left.Right.Right.Left = util.NewNode(9)
	t.Root.Left.Right.Right.Right = util.NewNode(10)

	t.Root.Right.Right = util.NewNode(6)
	t.PrintTree()

	target := t.Root.Left.Right
	k := 2

	res := make([]interface{}, 0)
	printNodesFromKthDistance(t.Root, target, k, &res)
	sort.Slice(res, func(i, j int) bool {
		return res[i].(int) < res[j].(int)
	})

	fmt.Printf("The nodes from %dth distance of given target %d are ", k, target.Data)
	for _, v := range res {
		fmt.Print(v, " ")
	}
	fmt.Println()

}
