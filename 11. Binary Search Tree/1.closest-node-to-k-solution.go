package main

import (
	"./util"
	"fmt"
	"math"
)

var t util.Tree

func closestNodeToK(k interface{}) interface{} {
	temp := t.Root
	minDiff := math.MaxInt64
	var num interface{}
	for temp != nil {
		currDiff := int(math.Abs(float64(temp.Data.(int) - k.(int))))

		if currDiff == 0 {
			return temp.Data
		}

		if currDiff < minDiff {
			minDiff = currDiff
			num = temp.Data
		}
		if temp.Data.(int) > k.(int) {
			temp = temp.Left
		} else {
			temp = temp.Right
		}
	}

	return num
}

func main() {
	t.Root = t.CreateNode(8)
	t.Root.Left = t.CreateNode(3)
	t.Root.Right = t.CreateNode(10)

	t.Root.Left.Left = t.CreateNode(1)
	t.Root.Left.Right = t.CreateNode(6)
	t.Root.Left.Right = t.CreateNode(6)
	t.Root.Left.Right.Left = t.CreateNode(4)
	t.Root.Left.Right.Right = t.CreateNode(7)

	t.Root.Right.Right = t.CreateNode(14)
	t.Root.Right.Right.Left = t.CreateNode(13)
	t.PrintTree()

	k := 16
	result := closestNodeToK(k)

	fmt.Printf("The closest element to %d is %d \n", k, result)
}
