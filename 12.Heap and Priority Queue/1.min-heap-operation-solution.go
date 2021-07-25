package main

import (
	"./util"
	"container/heap"
	"fmt"
)

func main() {
	mheap := make(util.MinHeap, 0)
	heap.Init(&mheap)

	heap.Push(&mheap, 1)
	heap.Push(&mheap, 8)
	heap.Push(&mheap, 3)
	heap.Push(&mheap, 6)
	heap.Push(&mheap, 2)
	heap.Push(&mheap, 7)
	heap.Push(&mheap, 5)

	for mheap.Len() > 0 {
		fmt.Print(heap.Pop(&mheap), " ")
	}
	fmt.Println()
}
