package main

import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 不用保持原来顺序，最后一个元素覆盖被删除的元素
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	stack := []int{5, 6, 7, 8}
	fmt.Println(stack)
	stack = append(stack, 9) // push v
	fmt.Println(stack)
	top := stack[len(stack)-1] // top of stack
	fmt.Println(top)
	fmt.Println(stack)
	stack = stack[:len(stack)-1] // pop
	fmt.Println(stack)
	stack = remove(stack, 2)
	fmt.Println(stack)
	stack = remove2(stack, 1)
	fmt.Println(stack)
}
