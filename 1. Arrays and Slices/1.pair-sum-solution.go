package main

import (
	"fmt"
)

func pairSum(arr []int, target int) []int {
	pa := make([]int, 0, 2)
	set := make(map[int]struct{})
	for _, v := range arr {
		if _, ok := set[target-v]; ok {
			pa = append(pa, target-v)
			pa = append(pa, v)
			return pa
		} else {
			set[v] = struct{}{}
		}
	}

	return pa
}

func main() {
	arr := []int{10, 5, 2, 3, -6, 9, 11}
	target := 4

	result := pairSum(arr, target)

	if len(result) == 0 {
		fmt.Println("The pair doesn't exist")
	} else {
		fmt.Printf("The Pair is [%d, %d] \n", result[0], result[1])
	}
}
