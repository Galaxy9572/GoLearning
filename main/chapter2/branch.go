package main

import "fmt"

func ifFunc(data int) bool {
	if data < 0 {
		return true
	} else if data == 0{
		return false
	}
	return true
}

func switchFunc(data int) {
	switch data {
	case 0:
		// 默认会有break
		fmt.Print("0")
	case 1:
		fmt.Print("1")
	default:
		// 报错
		panic("error")
	}
}
