package main 

import(
	"fmt"
	"sync"
	"sync/atomic"
) 

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全
//特定情况下比如简单的数值类型操作，累加， 无锁性能更强
func Add(num int) int {
	return num+1
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	num := 0
	for i:=0;i<10;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for j:=0;j<1000;j++{
				//在实际递增的位置加锁，性能更好，不要在外层，不然相当于 每个routines是串行执行的
				m.Lock()
				num = Add(num)
				m.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Println("num is :", num)

	num2 := int32(0)
	for i:=0;i<10;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for j:=0;j<1000;j++{
				atomic.AddInt32(&num2, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("num2 is :", num2)
}