package chapter10

import (
	"fmt"
)

// 协程：轻量级线程
// 非抢占式多任务处理，由协程主动交出控制权
// 编译器/解释器/虚拟机层面的多任务
// 一个线程可能有多个协程
// main退出，其他协程就会结束
// go run -race xx.go可以查看竞态条件

// goroutine可能切换的点I/O，select
// channel
// 等待锁
// 函数调用
// runtime.Gosched()
func ConcurrencyRun() {
	for i := 0; i < 10; i++ {
		// go func会使用协程
		go func(n int) {
			fmt.Printf("I am Coroutine %d\n", n)
			// 此方法会使协程交出控制权
			//runtime.Gosched()
		}(i)
	}
}
