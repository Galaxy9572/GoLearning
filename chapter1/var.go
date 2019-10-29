package chapter1

import "fmt"

func variables() {

	// 基本定义方式 ，定义的变量不使用编译器会报错
	var num int = 1
	fmt.Print(num)

	// 也可以使用类型推断
	var num2 = 2
	fmt.Print(num2)

	// 更可以使用:=赋初值
	num3 := 3
	fmt.Print(num3)

	// 可以连续给多个变量赋值
	var a, s, d = 3, 4, true
	fmt.Println(a, s, d)

	// 同样可以用:=给多个变量赋值
	z, x, c := 2, 3, 4
	fmt.Println(z, x, c)

}
