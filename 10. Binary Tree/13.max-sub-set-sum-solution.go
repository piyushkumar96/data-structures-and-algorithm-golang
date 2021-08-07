package main

import (
	"./util"
	"fmt"
	"math"
)

type IEPair struct {
	inc int
	exc int
}

func maxSubSetSUm(root *util.Node) IEPair {
	var iep IEPair
	if root == nil {
		iep.inc = 0
		iep.exc = 0
		return iep
	}

	lpa := maxSubSetSUm(root.Left)
	rpa := maxSubSetSUm(root.Right)

	iep.inc = root.Data.(int) + lpa.exc + rpa.exc
	iep.exc = int(math.Max(float64(lpa.inc), float64(lpa.exc))) + int(math.Max(float64(rpa.inc), float64(rpa.exc)))
	return iep
}

func main() {
	var t util.Tree
	t.Root = t.BuildTree()
	t.PrintTree()
	pa := maxSubSetSUm(t.Root)

	fmt.Println("\nThe maximum sum is ", int(math.Max(float64(pa.inc), float64(pa.exc))))

}
