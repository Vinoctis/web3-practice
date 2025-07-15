package main 

import (
	"sync"
	"fmt"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。


// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

// 接收端用for range或val, ok := <-ch判断通道状态
// 只有发送者才能关闭通道
// 已关闭的通道不能再发送数据

func Producer(c chan int, len int) {
	for i:=1;i<=len;i++{
		c <- i
	}
	close(c)
}

func Consumer(c chan int) {
	for i:= range c {
		// 使用range自动检测通道关闭
		fmt.Println("消费者协程接收数据：",i)
	}
}

func main () {
	var wg sync.WaitGroup

	c := make(chan int)
	wg.Add(2)
	go func(){
		defer wg.Done()
		Producer(c, 10)
	}()
	go func(){
		defer wg.Done()
		Consumer(c)
	}()
	wg.Wait()
	fmt.Println("任务1执行完成")

	wg.Add(2)
	c2 := make(chan int, 20)
	go func(){
		defer wg.Done()
		Producer(c2, 20)
	}()
	go func(){
		defer wg.Done()
		Consumer(c2)
	}()
	wg.Wait()
	fmt.Println("任务2执行完成")
}