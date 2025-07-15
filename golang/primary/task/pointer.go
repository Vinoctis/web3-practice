package main

//修改指针指向的数指
func Add(p_num *int) *int {
	*p_num += 10 
	return p_num
}

func main() {
	num := 100
	p_num := &num
	println(*p_num)
	println(Add(p_num))
}