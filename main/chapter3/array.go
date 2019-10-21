package main

import "fmt"

// 数组是值类型，当做参数传递会拷贝一份

func array() {
	// 定义常规数组
	var arr1 [5]int
	// 定义数组并初始化
	arr2 := [3]int{1, 2, 3}
	modifyArray(arr2)
	modifyArray2(&arr2)
	// 定义slice
	arr3 := [...]int{1, 2, 3, 4, 5}
	// 定义二维数组
	var twoDimArr [3][5]int

	// 遍历打印下标和值
	for i, v := range arr3 {
		fmt.Print(i, v)
	}

	fmt.Print(arr1, arr2, arr3, twoDimArr)
}

func modifyArray(arr [3]int) {
	// 并不会改到arr2的值，因为是值传递
	arr[0] = 100
}

func modifyArray2(arr *[3]int) {
	// 会改到arr2的值，因为是指针传递
	arr[0] = 100
}
