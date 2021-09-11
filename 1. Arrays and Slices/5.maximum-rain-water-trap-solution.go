package main

import (
	"fmt"
	"math"
)

func maxWaterTrapped(height []int, n int) int {
	left := make([]int, n)
	right := make([]int, n)

	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = int(math.Max(float64(left[i-1]), float64(height[i])))
	}

	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = int(math.Max(float64(right[i+1]), float64(height[i])))
	}

	water := 0
	for i := 0; i < n; i++ {
		water += int(math.Min(float64(left[i]), float64(right[i]))) - height[i]
	}

	return water
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	waterTrapped := maxWaterTrapped(height, len(height))
	fmt.Print("The maximum water trapped ", waterTrapped)

}
