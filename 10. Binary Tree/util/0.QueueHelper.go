package util

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

type Queue []*Node

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(item *Node) {
	*q = append(*q, item)
}

func (q *Queue) Pop() (*Node, bool) {
	if q.IsEmpty() {
		return nil, false
	} else {
		ele := (*q)[0]
		*q = (*q)[1:]
		return ele, true
	}
}

func (q *Queue) peek() (*Node, bool) {
	if q.IsEmpty() {
		return nil, false
	} else {
		return (*q)[0], true
	}
}
