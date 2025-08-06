package main 

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

	slow := 0 
	for fast:=1; fast<len(nums);fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

func main() {
	nums := []int{1,1,2,2,3,4,4,4,5,5,5,6}
	length := removeDuplicates(nums)
	fmt.Println(nums[:length]) // 输出: [1 2 3 4 5]
}