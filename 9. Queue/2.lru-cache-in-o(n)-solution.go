package main

import "fmt"

type LRUCache struct {
	capacity int
	list     []interface{}
	mp       map[interface{}]int
}

func Init(size int) *LRUCache {
	return &LRUCache{
		capacity: size,
		list:     make([]interface{}, 0),
		mp:       make(map[interface{}]int),
	}
}

func (l *LRUCache) refer(x interface{}) {
	if _, ok := l.mp[x]; !ok {
		if len(l.list) == l.capacity {
			delete(l.mp, l.list[0])
			l.list = l.list[1:]
		}
	} else {
		i := 0
		for ; i < len(l.list); i++ {
			if l.list[i] == x {
				break
			}
		}
		//fmt.Println(x, i)

		l.list = append(l.list[0:i], l.list[i+1:]...)
	}

	l.list = append(l.list, x)
	l.mp[x] = len(l.list)
}

func (l *LRUCache) printCache() {
	for _, v := range l.list {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	lru := Init(4)

	pages := []int{1, 2, 3, 1, 2, 4, 5, 2, 5, 1, 1}
	for _, v := range pages {
		lru.refer(v)
	}

	lru.printCache()
}
