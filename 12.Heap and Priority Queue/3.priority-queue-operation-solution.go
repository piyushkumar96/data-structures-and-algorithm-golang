package main

import (
	"./util"
	"container/heap"
	"fmt"
)

func main() {
	pq := make(util.PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &util.Item{Value: 0})
	heap.Push(&pq, &util.Item{Value: 1, Priority: 20, Index: 1})
	heap.Push(&pq, &util.Item{Value: 2, Priority: 10, Index: 2})
	heap.Push(&pq, &util.Item{Value: 3, Priority: 2, Index: 3})

	fmt.Println("The Priority Queue")
	for pq.Len() > 0 {
		item := (heap.Pop(&pq)).(*util.Item)
		fmt.Println(item.Value, item.Priority, item.Index)
	}

}
