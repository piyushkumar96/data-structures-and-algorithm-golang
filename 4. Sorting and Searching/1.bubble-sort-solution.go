package main

import (
	"fmt"
)

func bubbleSort(arr *[]int) {
	flag := false
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j+1] < (*arr)[j] {
				(*arr)[j+1], (*arr)[j] = (*arr)[j], (*arr)[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func main() {
	arr := []int{6, 3, 7, 1, 5, 8, 9, 2, 5}

	fmt.Println("Unsorted array ")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
	bubbleSort(&arr)
	fmt.Println("Sorted array ")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Print()
}
