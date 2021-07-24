package main

import "fmt"

func stringSubSets(s string, ind int, curr string, res *[]string) {
	if len(s) == ind {
		*res = append(*res, curr)
		return
	}

	stringSubSets(s, ind+1, curr+s[ind:ind+1], res)
	stringSubSets(s, ind+1, curr, res)
}

func main() {
	str := "abc"

	res := make([]string, 0)

	stringSubSets(str, 0, "", &res)

	fmt.Println("The subsets are ")
	for _, v := range res {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
