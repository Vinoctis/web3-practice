package main 

import(
	"fmt"
)

//题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
//[1,2,3] -> [1,2,4] , [9,9] -> [1,0,0]
//思路：从后往前遍历，加一，判断是否需要进位，
//不需要进位则直接返回，需要进位则继续往前遍历，直到不需要进位，或遍历到第一个元素，需要进位则在数组头部插入1


func PlusOne(s []int) []int {
	for i:=len(s)-1;i>=0;i-- {
		if s[i] < 9 {
			s[i]++
			return s
		}
		s[i] = 0
	}
	return append([]int{1}, s...)
}

func main() {
	fmt.Println(PlusOne([]int{1,2,3}))  // [1 2 4]
    fmt.Println(PlusOne([]int{9,9,9}))  // [1 0 0 0]
    fmt.Println(PlusOne([]int{9}))      // [1 0]
}