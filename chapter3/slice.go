package chapter3

import "fmt"

// slice实际上是数组的一个视图，个人认为可以理解成某个数组的引用 ？
func slice() {

	// arr := []int{0,1,2,3,4,5,6,7,8}同样可以
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//实际上是[2,6)，前闭后开区间
	s := arr[2:6]

	// [2,3,4,5]
	fmt.Print(s)

	s1 := arr[:6]

	s2 := arr[2:]

	s3 := arr[:]
	// [0,1,2,3,4,5]
	fmt.Print(s1)
	// [2,3,4,5,6,7,8]
	fmt.Print(s2)
	// [0,1,2,3,4,5,6,7,8]
	fmt.Print(s3)

	updateSlice(s2)

	//2,3,100,5,6,7,8
	fmt.Print(s2)
	// 2,3,4,5
	s4 := arr[2:6]
	// 5,6
	s5 := s4[3:5]
	// 不会报错，slice可以扩展，可以读到arr的下标，但只能看到后面的
	fmt.Print(s5)
	// 2
	fmt.Print(len(s5))
	// 4，s5的capacity向后可以扩展到s4的长度
	fmt.Print(cap(s5))

	// 如果append超过s5的cap，会分配一个新的底层数组
	s5 = append(s5, 10)

	// 创建一个len为6，cap为12的slice，slice扩容的方式为翻倍
	s6 := make([]int, 6, 12)
	// 从s5复制元素到s6
	copy(s6, s5)

}

func updateSlice(slice []int) {
	slice[2] = 100
}
