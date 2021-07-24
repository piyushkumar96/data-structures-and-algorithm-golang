package main

import (
	"fmt"
)

func longestBandLength(arr []int) int {
	maxLen := 0
	set := make(map[int]struct{}, 0)

	for i := 0; i < len(arr); i++ {

		count := 1
		tmp := arr[i] - 1
		_, ok := set[tmp]
		for ok {
			count++
			tmp--
			_, ok = set[tmp]
		}
		tmp = arr[i] + 1
		_, ok = set[tmp]
		for ok {
			count++
			tmp++
			_, ok = set[tmp]
		}
		set[arr[i]] = struct{}{}
		if maxLen < count {
			maxLen = count
		}
	}

	return maxLen
}

func main() {
	arr := []int{1, 11, 9, 3, 0, 18, 5, 2, 4, 10, 7, 12, 6, 8}

	res := longestBandLength(arr)
	fmt.Println("The max length mountain is ", res)
}
