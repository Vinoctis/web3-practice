package main

import (
	"strconv"
	"fmt"

)
func isPalindrome(num int) bool {

	str := strconv.Itoa(num)
	runes := []rune(str)
	for i, j := 0, len(str) - 1 ; i < j;i,j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return str == string(runes)
}

func isPalindromeByNumeric(num int) bool {
	// 0 和能被10整除都不是回文数
	if num == 0 || (num%10 == 0 && num != 0) {
		fmt.Println(num,"不是回文数")
	}

	reserve := 0 
	for num > reserve {
		reserve = reserve * 10 + num%10
		num /= 10 
	}
	
	return num == reserve || num == reserve/10
}

func main() {
	nums := []int{12321, 123, 45654, 1}
	for _, v := range nums {
		//字符串方法
		if isPalindrome(v) {
			fmt.Println(v,"是回文数")
		} else {
			fmt.Println(v,"不是回文数")
		}
		//数组方法
		if isPalindromeByNumeric(v) {
			fmt.Println(v,"是回文数")
		} else {
			fmt.Println(v,"不是回文数")
		}
	}
}