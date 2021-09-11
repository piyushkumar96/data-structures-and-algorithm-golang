package main

import (
	"fmt"
	"math"
)

func maxSubArraySum(arr []int, n int) int {
	maxSum := math.MinInt64
	currSum := 0

	for i := 0; i < n; i++ {
		currSum = int(math.Max(float64(arr[i]), float64(currSum+arr[i])))
		maxSum = int(math.Max(float64(maxSum), float64(currSum)))
	}

	if maxSum == math.MinInt64 {
		return -1
	}
	return maxSum
}

func main() {
	arr := []int{-1, 2, 3, 4, -2, 6, -8, 3}

	res := maxSubArraySum(arr, len(arr))
	fmt.Print("The sub array ", res)
}
