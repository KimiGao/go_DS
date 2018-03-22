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
func (tree *BiTree) add(t BiTree) {
	if tree.data > t.data {
		if tree.left == nil {
			tree.left = &t
		} else {
			tree.left.add(t)
		}
	} else {
		if tree.right == nil {
			tree.right = &t
		} else {
			tree.right.add(t)
		}
	}
}

// 中序遍历
func (tree *BiTree) travel() {
	if tree.left != nil {
		tree.left.travel()
	}

	fmt.Println(tree.data)

	if tree.right != nil {
		tree.right.travel()
	}
}
