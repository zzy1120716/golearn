package main

import (
	"fmt"
	"golearn/tutorial/5-2/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	//q.Push("abc")
	//fmt.Println(q.Pop())
	// panic: interface conversion: interface {} is string, not int

}
