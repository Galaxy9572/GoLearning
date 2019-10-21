package main

import (
	"fmt"
	"math/cmplx"
)

func complex() {
	// 复数定义
	var complex = 3 + 4i
	var res = cmplx.Abs(complex)
	fmt.Println(res)

}
