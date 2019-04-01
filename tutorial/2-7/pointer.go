package main

import "fmt"

// 交换两个变量的值
// 由于是值传递，所以必须传入指针类型
func swap1(a, b *int) {
	*a, *b = *b, *a
}

// 通过多返回值交换，调用时用两个int接收
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	a, b := 3, 4
	swap1(&a, &b)
	fmt.Println(a, b)

	a, b = swap2(a, b)
	fmt.Println(a, b)
}
