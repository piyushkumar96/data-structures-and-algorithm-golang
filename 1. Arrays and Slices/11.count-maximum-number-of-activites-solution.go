package main

import (
	"fmt"
	"sort"
)

func countActivities(arr [][2]int, n int) int {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] != arr[j][1] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})

	count := 1
	j := 0
	for i := 1; i < n; i++ {
		if arr[i][0] >= arr[j][1] {
			count++
			j = i
		}
	}

	return count
}

func main() {
	arr := [][2]int{{7, 9}, {0, 10}, {4, 5}, {8, 9}, {4, 10}, {5, 7}, {10, 11}}

	res := countActivities(arr, len(arr))
	fmt.Print("The max number of activities are ", res)
}
