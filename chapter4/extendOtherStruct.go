package chapter4

// 扩充别的类型
// 方法1：使用组合
// 方法2：定义别名

// 使用组合
type extendTreeNode struct {
	node *TreeNode
}

func (extendNode *extendTreeNode) postOrder() {
	if extendNode == nil || extendNode.node == nil {
		return
	}

	left := extendTreeNode{extendNode.node.Left}
	right := extendTreeNode{extendNode.node.Right}

	left.postOrder()
	right.postOrder()
	extendNode.node.Say()
}

// 使用别名
type Queue []int

func (q *Queue) push(v int) {
	*q = append(*q, v)
}

func (q *Queue) pop() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}
