package chapter8

import "fmt"

func PanicFunc() {
	// panic作用
	// 停止当前函数执行
	// 一直向上返回,执行每一层的defer语句
	// 如果没有遇见recover，程序退出
	panic("Error!!!!!!")
}

func RecoverFunc() {
	// 可以获取panic的值
	// 仅在defer中使用
	// 无法处理可以重新panic

	defer func() {
		i := recover()
		if i != nil {
			fmt.Print("Error Caught")
		} else {
			panic("Error Again")
		}
	}()

	panic("Error!!!!!!")
}
