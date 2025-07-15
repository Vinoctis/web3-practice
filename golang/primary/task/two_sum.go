package main 

//题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		complement := target -  num 
		if _, ok := m[complement]; ok {
			return []int{complement, num}
		} 
		m[num] = i
	}
	return nil
}

func main() {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
}
