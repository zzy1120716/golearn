// recover的使用，结合panic
package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover()
		// 类型断言
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			//panic(r)
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()

	// panic(errors.New("this is an error"))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	// 不知道错误含义的情况下，进行re-panic
	panic(123)
}

func main() {
	tryRecover()
}
