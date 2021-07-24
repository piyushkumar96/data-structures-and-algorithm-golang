package main

import (
	"fmt"
	"strings"
)

func stringTokenization(s string) []string {
	return strings.Split(s, " ")
}

func main() {
	str := "My name is Piyush Kumar"

	resTok := stringTokenization(str)

	for _, v := range resTok {
		fmt.Print(v, ", ")
	}
	fmt.Println()
}
