package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 为结构体定义方法
// 函数名前的括号里，称为接收者
// 值接收者
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

// 指针接收者
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " +
			"node. Ignored.")
		// 此处需要return，否则无法从nil中拿到value，会报错
		return
	}
	node.value = value
}

// 遍历二叉树
func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 使用自定义工厂函数，实质是普通函数，
// 作用等价于其他语言中的构造函数。
// 注意返回了局部变量的地址！
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	fmt.Println(root)

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	// new是内建方法，用来创建结构体实例
	// go语言中，不论地址还是结构本身，都用“.”访问成员
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	//fmt.Println(root)

	/*nodes := []treeNode {
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)*/

	root.print()
	fmt.Println()
	root.right.left.setValue(4)
	/*root.right.left.print()
	fmt.Println()

	root.print()
	root.setValue(100)
	fmt.Println()

	pRoot := &root
	pRoot.print()
	fmt.Println()
	pRoot.setValue(200)
	pRoot.print()
	fmt.Println()

	var nRoot *treeNode
	nRoot.setValue(200)
	nRoot = &root
	nRoot.setValue(300)
	nRoot.print()*/

	root.traverse()
}
