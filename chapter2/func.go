package chapter2

import (
	"fmt"
	"strconv"
)

func oneReturn(a int, b float64) {
	fmt.Print(a, b)
}

// 返回两个值的函数定义
func twoReturn(a, b int) (int, string) {
	// strconv.Itoa()把数字转为字符串
	return a + b, strconv.Itoa(a + b)
}

// 返回值可以取名
func funcReturnName(a, b int) (r1 int, r2 string) {

	// _用来不接收返回值
	q, _ := twoReturn(1, 2)
	fmt.Print(q)
	// strconv.Itoa()把数字转为字符串
	return a + b, strconv.Itoa(a + b)
}

// 函数式编程
func funcFunc(fun func(int, int) int, c, d int) int {
	return fun(c, d) + c*d
}

// 可变参数列表
func multipleArgsFunc(args ...int) {

	for i := range args {
		fmt.Print(i)
		fmt.Print(args[i])
	}
}
