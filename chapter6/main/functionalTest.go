package main

import (
	"GoLearning/chapter6"
	"fmt"
)

func main() {
	fun := chapter6.Add()
	for i := 1; i < 10; i++ {
		i2 := fun(i)
		fmt.Println(i2)
	}

	fibonacci := chapter6.Fibonacci()
	i := fibonacci(5)
	fmt.Print(i)
}
