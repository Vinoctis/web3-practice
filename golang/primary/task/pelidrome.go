package main

import (
	"strconv"
	"fmt"

)
func isPalindrome(num int) {

	str := strconv.Itoa(num)
	runes := []rune(str)
	for i, j := 0, len(str) - 1 ; i < j;i,j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	if str == string(runes) {
		fmt.Println(num,"是回文数")
	} else {
		fmt.Println(num,"不是回文数")
	}
}

func main() {
	isPalindrome(123321)
	isPalindrome(123)
	isPalindrome(45654)
	isPalindrome(1)
}