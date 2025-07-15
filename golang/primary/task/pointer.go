package main

import(
	"fmt"
)

//题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func Add(p_num *int) *int {
	*p_num += 10 
	fmt.Println("new num is :",*p_num)
	return p_num
}

//题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
//for _, v := range *s 中的v是切片元素的副本，对其修改不会影响原数据
// 即使传递的是切片指针，也需要通过索引访问元素才能修改原值
// 切片本身是引用类型，通常不需要传递指针（除非需要修改切片长度）
func Multiply(s []int) []int {
	for i := range s {
		s[i] *= 2
	}
	return s
}

func main() {
	num := 100
	p_num := &num
	fmt.Println("old num is ", *p_num)
	Add(p_num)

	num2 := []int{2,3,4} //切片是引用传递
	fmt.Println("old slice is :", num2)
	fmt.Println("new slice is :", Multiply(num2))
}