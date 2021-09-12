package main

import (
	"fmt"
)

func arrayProducts(arr []int, n int) []int {
	res := make([]int, n)

	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * arr[i-1]
	}

	right := make([]int, n)
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * arr[i+1]
	}

	for i := 0; i < n; i++ {
		res[i] = left[i] * right[i]
	}
	return res
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	res := arrayProducts(arr, len(arr))
	fmt.Print("The resultant product array is", res)
}
