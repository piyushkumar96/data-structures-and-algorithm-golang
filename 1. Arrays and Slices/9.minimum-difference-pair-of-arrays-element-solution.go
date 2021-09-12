package main

import (
	"fmt"
	"math"
	"sort"
)

func minDifferenceElement(arr1 []int, n int, arr2 []int, m int) [2]int {

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	i, j := 0, 0
	minDiff := math.Abs(float64(arr1[0] - arr2[0]))
	pair := [2]int{arr1[0], arr2[0]}
	for i < n && j < m {
		if minDiff > math.Abs(float64(arr1[i]-arr2[j])) {
			minDiff = math.Abs(float64(arr1[i] - arr2[j]))
			if arr1[i] > arr2[j] {
				pair = [2]int{arr2[j], arr1[i]}
			} else {
				pair = [2]int{arr1[i], arr2[j]}
			}
		}

		if arr1[i] > arr2[j] {
			j++
		} else {
			i++
		}
	}

	return pair
}

func main() {
	arr1 := []int{23, 5, 10, 17, 30}
	arr2 := []int{26, 134, 135, 14, 19}

	minDiff := minDifferenceElement(arr1, len(arr1), arr2, len(arr2))
	fmt.Print("The minimum diff is between ", minDiff)
}
