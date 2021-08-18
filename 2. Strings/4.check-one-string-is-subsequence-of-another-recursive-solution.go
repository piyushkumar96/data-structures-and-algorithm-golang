package main

import "fmt"

func checkSubsequenceRecursive(s1 string, l1 int, s2 string, l2 int) bool {
	if l1 == 0 {
		return false
	}
	if l2 == 0 {
		return true
	}

	if s1[l1] == s2[l2] {
		return checkSubsequenceRecursive(s1, l1-1, s2, l2-1)
	}

	return checkSubsequenceRecursive(s1, l1-1, s2, l2)
}

func main() {
	var s1, s2 string
	fmt.Print("Enter string s1 ")
	_, _ = fmt.Scan(&s1)

	fmt.Print("Enter string s2 ")
	_, _ = fmt.Scan(&s2)

	res := checkSubsequenceRecursive(s1, len(s1)-1, s2, len(s2)-1)
	if res {
		fmt.Println("The s2 is subsequence of s1")
	} else {
		fmt.Println("The s2 is not subsequence of s1")
	}
}
