package util

import (
	"fmt"
	"math"
)

type Tree struct {
	Root *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (t *Tree) BuildTree() *Node {
	var item int
	_, _ = fmt.Scan(&item)

	if item == -1 {
		return nil
	}

	node := &Node{
		Data:  item,
		Left:  nil,
		Right: nil,
	}

	node.Left = t.BuildTree()
	node.Right = t.BuildTree()

	return node
}
func LevelOrderTraversal(root *Node) {
	var q Queue
	q.Push(root)
	q.Push(nil)
	for !q.IsEmpty() {
		node, _ := q.Pop()
		if node == nil {
			fmt.Println()
			if !q.IsEmpty() {
				q.Push(nil)
			}
		} else {
			fmt.Print(node.Data, " ")
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}
}

func (t *Tree) PrintTree() {
	if t.Root == nil {
		fmt.Println("Tree is empty")
		return
	}
	fmt.Println("Tree: ")
	LevelOrderTraversal(t.Root)
}

func HeightOfTree(root *Node) int {
	if root == nil {
		return 0
	}

	lh := HeightOfTree(root.Left)
	rh := HeightOfTree(root.Right)

	return int(math.Max(float64(lh), float64(rh))) + 1
}
