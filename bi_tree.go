package main

import "fmt"

type BiTree struct {
	data  int
	left  *BiTree
	right *BiTree
}

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

func main() {
	tree := BiTree{data: 13}

	tree.add(BiTree{data: 4})
	tree.add(BiTree{data: 2})
	tree.add(BiTree{data: 14})
	tree.add(BiTree{data: 1})
	tree.add(BiTree{data: 9})
	tree.add(BiTree{data: 21})
	tree.add(BiTree{data: 11})
	tree.add(BiTree{data: 8})
	tree.add(BiTree{data: 7})
	tree.add(BiTree{data: 24})

	tree.travel()
}
