package chapter6

// 函数式编程，返回函数
func Add() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func Fibonacci() func(time int) int {
	a, b := 0, 1
	return func(time int) int {
		a, b = b, a+b
		return a
	}
}
