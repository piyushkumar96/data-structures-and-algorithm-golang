package util

type MaxHeap []int

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MaxHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *MaxHeap) Pop() interface{} {
	n := len(*mh)
	ele := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return ele
}
