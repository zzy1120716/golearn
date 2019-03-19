package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	//ch := make(chan int, 2)		// 输出变为不确定
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}
