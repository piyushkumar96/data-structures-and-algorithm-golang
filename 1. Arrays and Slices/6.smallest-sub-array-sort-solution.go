package main

import (
	"fmt"
)

func smallestSubArraySort(arr []int, n int) [2]int {
	var s int
	res := [2]int{-1, -1}
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			s = i
			break
		}
	}

	if s == n-1 {
		return res
	}

	var e int
	for i := n - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			e = i
			break
		}
	}

	min := arr[s]
	max := arr[s]
	for i := s + 1; i <= e; i++ {
		if min > arr[i] {
			min = arr[i]
		}

		if max < arr[i] {
			max = arr[i]
		}
	}

	for i := 0; i < s; i++ {
		if min < arr[i] {
			s = i
			break
		}
	}

	for i := n - 1; i >= e+1; i-- {
		if max > arr[i] {
			e = i
			break
		}
	}

	return [2]int{s, e}
}

func main() {
	arr := []int{10, 12, 20, 30, 25, 40, 32, 31, 35, 50, 60}

	res := smallestSubArraySort(arr, len(arr))
	fmt.Print("The sub array ", res)

}
