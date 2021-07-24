package main

import (
	"fmt"
)

func highestLengthMountain(arr []int) int {
	maxLen := 1
	n := len(arr)
	for i := 0; i < n-1; i++ {

		if arr[i] > arr[i+1] {
			j, k := i, i
			for j > 0 && arr[j] > arr[j-1] {
				j--
			}
			for k < n && arr[k] > arr[k+1] {
				k++
			}

			if maxLen < (k-j)+1 {
				maxLen = k - j + 1
			}
			i = k
		}
	}

	return maxLen
}

func main() {
	arr := []int{5, 6, 1, 2, 3, 4, 5, 4, 3, 2, 0, 1, 2, 3, -2, 4}

	res := highestLengthMountain(arr)
	fmt.Println("The max length mountain is ", res)
}
