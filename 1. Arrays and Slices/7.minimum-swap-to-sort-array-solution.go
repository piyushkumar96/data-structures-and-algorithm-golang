package main

import (
	"fmt"
	"sort"
)

func minimumSwap(arr []int, n int) int {
	temp := make([]int, n)
	copy(temp, arr)

	mp := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		mp[arr[i]] = i
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})

	swapCount := 0
	for i := 0; i < n; i++ {
		if arr[i] != temp[i] {
			swapCount++
			init := arr[i]

			arr[i], arr[mp[temp[i]]] = arr[mp[temp[i]]], arr[i]
			mp[init] = mp[temp[i]]
			mp[temp[i]] = i
		}
	}

	return swapCount
}

func main() {
	arr := []int{10, 12, 20, 30, 25, 40, 32, 31, 35, 50, 60}

	res := minimumSwap(arr, len(arr))
	fmt.Print("The sub array ", res)
}
