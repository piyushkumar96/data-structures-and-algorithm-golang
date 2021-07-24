package main

import (
	"fmt"
)

func space20(inputStr []byte) string {
	spCount := 0
	for _, v := range inputStr {
		if string(v) == " " {
			spCount++
		}
	}

	resSize := len(inputStr) + 2*spCount - 1
	inputStr = append(inputStr, make([]byte, 3*spCount)...)
	for i := len(inputStr) - 3*spCount - 1; i >= 0; i-- {
		if string(inputStr[i]) == " " {
			inputStr[resSize] = []byte("0")[0]
			inputStr[resSize-1] = []byte("2")[0]
			inputStr[resSize-2] = []byte("%")[0]
			resSize -= 2

		} else {
			inputStr[resSize] = inputStr[i]
		}
		resSize--
	}
	return string(inputStr)
}

func main() {
	charArr := make([]byte, 0, 10000)

	inputStr := " hello world, how are you ?"

	charArr = append(charArr, inputStr...)

	resultStr := space20(charArr)
	fmt.Printf("Given string is: \t %s \nResultant string is: \t %s \n", string(inputStr), resultStr)
}
