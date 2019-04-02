package main

import (
	"fmt"
	"golearn/tutorial/4-2/tree"
)

// 使用组合扩展类型
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	// 此处需要将tree.Node包装成myTreeNode，
	// 以使用myTreeNode结构的方法postOrder
	//myTreeNode{myNode.node.Left}.postOrder()
	//myTreeNode{myNode.node.Right}.postOrder()

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	fmt.Println(root)

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	// new是内建方法，用来创建结构体实例
	// go语言中，不论地址还是结构本身，都用“.”访问成员
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()
	fmt.Println()
	root.Right.Left.SetValue(4)
	root.Traverse()
	fmt.Println()

	// cannot call pointer method on myTreeNode literal
	// cannot take the address of myTreeNode literal
	//myTreeNode{&root}.postOrder()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
