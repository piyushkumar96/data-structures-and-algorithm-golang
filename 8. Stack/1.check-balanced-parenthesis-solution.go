package main

import (
	"fmt"
)

type Stack []interface{}

func (s *Stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}

func (s *Stack) push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) pop() (interface{}, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		ele := (*s)[index]
		*s = (*s)[:index]

		return ele, true
	}
}

func (s *Stack) peek() (interface{}, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		return (*s)[index], true
	}
}

func checkBalancedParenthesis(str string) bool {
	var st Stack

	for _, v := range str {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			st.push(string(v))
		} else if string(v) == ")" || string(v) == "}" || string(v) == "]" {
			switch string(v) {
			case ")":
				top, flag := st.pop()
				if !flag || top != "(" {
					return false
				}
				break
			case "}":
				top, flag := st.pop()
				if !flag || top != "{" {
					return false
				}
				break
			case "]":
				top, flag := st.pop()
				if !flag || top != "[" {
					return false
				}
				break
			default:
				return false

			}
		}
	}

	return st.isEmpty()
}

func main() {
	var str string
	fmt.Print("Enter the string for check ")
	_, _ = fmt.Scan(&str)

	if checkBalancedParenthesis(str) {
		fmt.Println("IsBalanced")
	} else {
		fmt.Println("NotBalanced")
	}
}
