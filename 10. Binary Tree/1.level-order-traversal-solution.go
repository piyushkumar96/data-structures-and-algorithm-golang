package main

import (
	"./util"
	"fmt"
)

type Tree struct {
	root *(util.Node)
}

func (t *Tree) buildTree() *(util.Node) {
	var item int
	_, _ = fmt.Scan(&item)

	if item == -1 {
		return nil
	}

	node := &util.Node{
		Data:  item,
		Left:  nil,
		Right: nil,
	}

	node.Left = t.buildTree()
	node.Right = t.buildTree()

	return node
}

func levelOrderTraversal(root *(util.Node)) {
	var q util.Queue
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

func (t *Tree) printTree() {
	if t.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	fmt.Println("Tree: ")
	levelOrderTraversal(t.root)
}

func main() {
	var t Tree
	t.root = t.buildTree()
	t.printTree()
}
