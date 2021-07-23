package main

import (
	"fmt"
)

type Queue []interface{}

func CreateQueue() *Queue {
	return &Queue{}
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) push(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) pop() (interface{}, bool) {
	if q.isEmpty() {
		return nil, false
	} else {
		ele := (*q)[0]
		*q = (*q)[1:]
		return ele, true
	}
}

func (q *Queue) popAtIndex(index int) bool {
	if q.isEmpty() {
		return false
	} else {
		*q = append((*q)[:index], (*q)[index+1:]...)
		return true
	}
}

func (q *Queue) peek() (interface{}, bool) {
	if q.isEmpty() {
		return nil, false
	} else {
		return (*q)[0], true
	}
}

func linearSearch(arr []interface{}, item string) int {
	for i, v := range arr {
		if item == v {
			return i
		}
	}

	return -1
}

func firstNonRepeatingLetter(str string) []string {
	q := CreateQueue()
	var res []string
	if len(str) == 0 {
		return res
	}
	m := make(map[string]int)

	for _, v := range str {
		_, ok := m[string(v)]
		if ok {
			index := linearSearch(*q, string(v))
			if index != -1 {
				q.popAtIndex(index)
			}
			if q.isEmpty() {
				res = append(res, "-1")

			} else {
				ele, _ := q.peek()
				res = append(res, ele.(string))
			}

		} else {
			m[string(v)] = 1
			if q.isEmpty() {
				res = append(res, string(v))

			} else {
				ele, _ := q.peek()
				res = append(res, ele.(string))
			}
			q.push(string(v))
		}
	}

	return res
}

func main() {
	var str string
	fmt.Print("Enter the string ")
	_, _ = fmt.Scan(&str)

	result := firstNonRepeatingLetter(str)

	for _, v := range result {
		fmt.Printf("%s ", string(v))
	}
	fmt.Println()

}
