package main

import "fmt"

func checkSubsequenceIterative(s1, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	j := 0

	for i := 0; i < l1 && j < l2; i++ {
		if s1[i] == s2[j] {
			j++
		}
	}

	return j == l2
}

func main() {
	var s1, s2 string
	fmt.Print("Enter string s1 ")
	_, _ = fmt.Scan(&s1)

	fmt.Print("Enter string s2 ")
	_, _ = fmt.Scan(&s2)

	res := checkSubsequenceIterative(s1, s2)
	if res {
		fmt.Println("The s2 is subsequence of s1")
	} else {
		fmt.Println("The s2 is not subsequence of s1")
	}
}
