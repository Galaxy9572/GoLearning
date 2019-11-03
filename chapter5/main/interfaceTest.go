package main

import (
	"HelloWorld/chapter5"
	"fmt"
)

func main() {
	var v chapter5.Vehicle
	v = chapter5.Car{"BMW"}
	v.Run()

	// 将v强转成Car类型
	car := v.(chapter5.Car)

	fmt.Println(car)

	// 打印类型和值
	fmt.Printf("%T %v", v, v)
}
