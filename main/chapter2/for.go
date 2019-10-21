package main

import (
	"bufio"
	"fmt"
	"os"
)

// golang没有while语句
func forFunc() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		// 死循环
		fmt.Print("I am Dead Lock")
	}

	file, err := os.Open("main/abc.txt")

	if err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

}
