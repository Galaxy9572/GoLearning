package main

import (
	"GoLearning/chapter7"
	"fmt"
)

func main() {
	fun := chapter7.Add()
	for i := 1; i < 10; i++ {
		i2 := fun(i)
		fmt.Println(i2)
	}

	fibonacci := chapter7.Fibonacci()
	i := fibonacci(5)
	fmt.Print(i)
}
