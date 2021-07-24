package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func getStringAtPosK(s string, k int) string {
	return strings.Split(s, " ")[k-1]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Scan(&n)

	strs := make([]string, 0)
	for i := 0; i < n; i++ {
		s, _ := in.ReadString('\n')
		strs = append(strs, s)
	}

	var pos int
	var reversal bool
	var cmpType string
	_, _ = fmt.Scanln(&pos, &reversal, &cmpType)

	keyList := make([][2]string, 0)

	for i := 0; i < n; i++ {
		key := getStringAtPosK(strs[i], pos)

		keyList = append(keyList, [2]string{
			strs[i], key,
		})
	}

	if cmpType == "lexicographic" {
		sort.Slice(keyList, func(i, j int) bool {
			return keyList[i][1] < keyList[j][1]
		})

	} else {
		sort.Slice(keyList, func(i, j int) bool {
			first, _ := strconv.Atoi(keyList[i][1])
			second, _ := strconv.Atoi(keyList[j][1])
			return first < second
		})
	}

	if reversal {
		swap := reflect.Swapper(keyList)
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			swap(i, j)
		}
	}

	fmt.Println("\nThe final list is:-")
	for _, v := range keyList {
		fmt.Print(v[0])
	}

	fmt.Println()
}
