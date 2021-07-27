package main

import (
	"./util"
	"fmt"
	"math"
)

type HDPair struct {
	height int
	dia    int
}

func diameterOfTree_ON(root *util.Node) HDPair {
	var hdp HDPair
	if root == nil {
		hdp.height = 0
		hdp.dia = 0
		return hdp
	}

	lpa := diameterOfTree_ON(root.Left)
	rpa := diameterOfTree_ON(root.Right)

	hdp.height = int(math.Max(float64(lpa.height), float64(rpa.height))) + 1
	d1 := lpa.height + rpa.height
	d2 := lpa.dia
	d3 := rpa.dia

	hdp.dia = int(math.Max(math.Max(float64(d1), float64(d2)), float64(d3)))
	return hdp
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	dia := diameterOfTree_ON(t.Root)
	fmt.Println("\nThe Diameter of Tree is ", dia)
}
