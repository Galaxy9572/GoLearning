package chapter4

import "fmt"

// 封装
// 名字用驼峰命名法
// 首字母大写为public
// 首字母小写为private

// 包：
// 每个目录只能有一个包，包名不一定和目录名一样
// main包包含可执行方法入口
// 为struct定义的函数必须在同一个包内，但可以是不同的文件内
func encapsulation() {

	node := TreeNode{}
	node.SetValue(3)
	fmt.Print(node)

}
