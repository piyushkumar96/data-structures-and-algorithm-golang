package main

import (
	"./util"
)

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
}
