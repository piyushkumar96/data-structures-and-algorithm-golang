package util

import "container/heap"

type Item struct {
	Value    interface{}
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	ele := (*pq)[n-1]
	ele.Index = -1
	(*pq)[n-1] = nil
	*pq = (*pq)[:n-1]

	return ele
}

func (pq *PriorityQueue) Update(item *Item, value int, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
