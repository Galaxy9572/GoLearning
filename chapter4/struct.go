package chapter4

import "fmt"

// go只支持封装，不支持继承和多态，没构造函数
type TreeNode struct {
	Left, Right *TreeNode
	Value       int
}

// 关于值接收者和指针接收者：
// 要改变结构体中的内容必须用指针接收者
// 结构过大要考虑指针接收者，因为值拷贝过多音响性能

// 结构体的函数定义方式，类似于Java的成员方法，是在函数名前括号内加入函数的接收者，自己本身
func (node TreeNode) Say() {
	fmt.Print(node.Value)
}

// 因为是值传递，所以赋值要传指针
func (node *TreeNode) SetValue(value int) {
	if node == nil {
		return
	}
	node.Value = value
}

// 数的中根序遍历
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Say()
	node.Left.Say()
	node.Right.Say()
}

func structFunc() {
	// 产生新结构体
	root := TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2, Left: nil, Right: nil}
	root.Right = &TreeNode{Value: 3, Left: nil, Right: nil}
	root.Right.Left = new(TreeNode)

	// 不需要写成&root.SetValue(x)
	root.SetValue(5)

	// nil指针也可以调用函数
	var nilNode *TreeNode
	nilNode.SetValue(1)

}
