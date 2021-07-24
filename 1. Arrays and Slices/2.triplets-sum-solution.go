package main

import (
	"fmt"
	"sort"
)

func tripletsSum(arr []int, sum int) [][3]int {
	res := make([][3]int, 0)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	asize := len(arr)
	for i := 0; i < asize; i++ {
		j := i + 1
		k := asize - 1

		for j < k {
			currSum := arr[i] + arr[j] + arr[k]
			if sum == currSum {
				res = append(res, [3]int{arr[i], arr[j], arr[k]})
				j++
				k--
			} else if sum < currSum {
				k--
			} else {
				j++
			}
		}
	}

	return res
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}
	sum := 18

	res := tripletsSum(arr, sum)

	if len(res) == 0 {
		fmt.Println("No triplets for given sum")
	} else {
		fmt.Println("Triplets :-")
		for _, v := range res {
			fmt.Println(v, " ")
		}
		fmt.Println()
	}
}
