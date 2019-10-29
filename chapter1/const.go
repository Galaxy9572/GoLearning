package chapter1

import "fmt"

// 函数外定义的变量为包内可见，无全局变量的概念
// 常量枚举定义
// iota为从1开始的自增序列
const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
	eb
)

func constants() {

	const c = "aaa"
	fmt.Print(b, kb, mb, gb, tb, pb, eb)

}
