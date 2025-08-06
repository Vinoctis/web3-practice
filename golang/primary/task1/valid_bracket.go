package main

import (
	"fmt"
)


func isValidBracket(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else if len(stack) > 0 && stack[len(stack)-1] == bracketMap[c] {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	s := "()[]{}"
	fmt.Println(isValidBracket(s))

	s = "([)]"
	fmt.Println(isValidBracket(s))

	s = "{[]}"
	fmt.Println(isValidBracket(s))
}
