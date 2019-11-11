package chapter8

import (
	"errors"
	"fmt"
)

func deferFunc() {
	fmt.Print("1")
	fmt.Print("2")

	// defer语句在return前执行
	// defer语句存在一个栈中，先进后出，所以这里先打印2，再打印1
	// 常用场景：关闭连接关闭流
	defer fmt.Print("defer 1")
	defer fmt.Print("defer 2")

	// defer伪代码
	// file.Read()
	// defer file.Close()
	// defer file.Flush()

	return

}

func createError() {
	e := errors.New("Error Message")
	if e != nil {
		fmt.Print(e)
	}
}
