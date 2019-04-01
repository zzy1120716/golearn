package main

import (
	"fmt"
	"golearn/tutorial/4-2/tree"
)

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
}
