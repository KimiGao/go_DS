package main

import "fmt"

type MyTree struct {
	data   string
	parent *MyTree
}

func (tree *MyTree) add(parent string, child string) {
	p_tree, c_tree = MyTree{}
}

func (tree *MyTree) getParent(name string) MyTree {

}

func (tree *MyTree) getChild(name string) []string {

}

func main() {
	tree := MyTree{data: "世界"}

	tree.add("世界", "亚洲")
	tree.add("世界", "欧洲")
	tree.add("世界", "美洲")
	tree.add("亚洲", "中国")
	tree.add("中国", "北京")
	tree.add("亚洲", "韩国")
	tree.add("亚洲", "日本")
	tree.add("欧洲", "英国")
	tree.add("欧洲", "法国")
	tree.add("美洲", "美国")
	tree.add("美洲", "阿根廷")

	tree.getParent("日本")
	tree.getParent("北京")

	tree.getChild("亚洲")
	tree.getChild("美洲")
}
