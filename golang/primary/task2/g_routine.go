package main 

import(
	"fmt"
	"time"
	"sync"
)
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func PrintEven() {
	for i:=2;i<=10;i+=2{
		fmt.Println("偶数：",i)
	}
}

func PrintOdd() {
	for i:=1;i<=10;i+=2{
		fmt.Println("奇数：",i)
	}
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

// 循环变量捕获：所有goroutine共享循环变量v，可能导致全部执行最后一个任务
// 竞态条件：当循环快速迭代时，goroutine可能读取到错误的v值
// 错误的任务执行：实际可能所有任务都执行同一个函数
// 修改说明：
// 将循环变量v作为参数显式传递给goroutine
// 每个goroutine会创建自己的task参数副本
// 保持原有功能的同时修复并发问题
// 测试验证： 当传入[]func(){PrintOdd, PrintEven}时，会正确输出两个任务的执行时间，而不是重复执行同一个任务。
func CalTime(tasks []func()) {
	fmt.Println("调度任务开始...")
	var wg sync.WaitGroup
	for _, v := range tasks {
		wg.Add(1)
		go func(t func()){
			defer wg.Done()
			begin := time.Now()
			t() 
			end := time.Now()
			fmt.Println("任务执行时间：", end.Sub(begin))
		}(v) 
	}
	wg.Wait()
	fmt.Println("所有任务执行完成")
}


func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func(){
		defer wg.Done()
		PrintOdd()
	}()

	go func(){
		defer wg.Done()
		PrintEven()
	}()

	wg.Wait()
	fmt.Println("done")

	tasks := []func(){
		PrintOdd,
		PrintEven,
	}

	CalTime(tasks)
}