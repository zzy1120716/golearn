package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 为结构体定义方法
// 函数名前的括号里，称为接收者
// 值接收者
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 指针接收者
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " +
			"node. Ignored.")
		// 此处需要return，否则无法从nil中拿到value，会报错
		return
	}
	node.Value = value
}

// 使用自定义工厂函数，实质是普通函数，
// 作用等价于其他语言中的构造函数。
// 注意返回了局部变量的地址！
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
