package main

import (
	"./util"
)

var t util.Tree

func sortedArrayToSelfBalanceBST(arr []int, s, e int) *util.Node {
	if s > e {
		return nil
	}
	mid := (s + e) / 2
	root := t.CreateNode(arr[mid])

	root.Left = sortedArrayToSelfBalanceBST(arr, s, mid-1)
	root.Right = sortedArrayToSelfBalanceBST(arr, mid+1, e)

	return root
}

func main() {
	sortedArr := []int{1, 2, 3, 4, 5, 6, 7}
	t.Root = sortedArrayToSelfBalanceBST(sortedArr, 0, len(sortedArr)-1)
	t.PrintTree()
}
