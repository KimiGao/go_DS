package ds

import "fmt"

type BiTree struct {
	data  int
	left  *BiTree
	right *BiTree
}

// 添加节点
// @param [*BiTree] 需要添加的节点
// @return [BiTree] 返回添加的节点
func (tree *BiTree) AddBT(t BiTree) {
	if tree.data > t.data {
		if tree.left == nil {
			tree.left = &t
		} else {
			tree.left.AddBT(t)
		}
	} else {
		if tree.right == nil {
			tree.right = &t
		} else {
			tree.right.AddBT(t)
		}
	}
}

// 中序遍历
func (tree *BiTree) Travel() {
	if tree.left != nil {
		tree.left.Travel()
	}

	fmt.Println(tree.data)

	if tree.right != nil {
		tree.right.Travel()
	}
}
