package main

import (
	"fmt"
)

func pointer() {

	a := 2
	// &a表示取a的内存地址
	// 也可以写成var pointerOfA *int = &a
	pointerOfA := &a
	// 指针值赋值为3
	*pointerOfA = 3
	// 那么a也会等于3
	fmt.Print(a)

}

// 交换两个值，golang都是值传递，可以通过指针实现引用传递
func swap(a, b *int) {
	*a, *b = *b, *a
}
