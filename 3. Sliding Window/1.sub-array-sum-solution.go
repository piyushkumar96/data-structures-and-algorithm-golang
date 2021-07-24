package main

import (
	"fmt"
)

func subArraySum(arr []int, k int) [][2]int {
	j, sum := 0, 0
	n := len(arr)
	res := make([][2]int, 0)
	for i, v := range arr {
		sum = sum + v
		for j < n && j <= i && sum > k {
			sum = sum - arr[j]
			j++
		}

		if sum == k {
			res = append(res, [2]int{j, i})
		}
	}

	return res
}

func main() {
	arr := []int{1, 3, 2, 1, 4, 1, 3, 2, 1, 1, 1}
	k := 8
	pairs := subArraySum(arr, k)
	if len(pairs) == 0 {
		fmt.Println("No pairs found")
	} else {
		fmt.Println("The pairs are")
		for _, v := range pairs {
			fmt.Printf("[%d, %d]\n", v[0], v[1])
		}
	}
	fmt.Println()

}
