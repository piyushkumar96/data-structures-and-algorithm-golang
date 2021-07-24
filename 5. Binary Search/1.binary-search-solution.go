package main

import (
	"fmt"
)

func binarySearch(arr []int, k int, s int, e int) int {
	if s > e {
		return -1
	}

	mid := (s + e) / 2

	if arr[mid] == k {
		return mid
	}
	if arr[mid] > k {
		return binarySearch(arr, k, s, mid-1)
	}

	return binarySearch(arr, k, mid+1, e)
}

func main() {
	arr := []int{1, 3, 8, 99, 100, 333, 500}
	k := 333

	res := binarySearch(arr, k, 0, len(arr)-1)
	if res == -1 {
		fmt.Println("The given element not found")
	} else {
		fmt.Println("The given element is found at index ", res)
	}
}
