package util

type Queue []interface{}

func CreateQueue() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	} else {
		ele := (*q)[0]
		*q = (*q)[1:]

		return ele, true
	}
}

func (q *Queue) Peek() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	} else {
		return (*q)[0], true
	}
}
