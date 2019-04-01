package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 增加error作为错误，返回
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	// 不接受第二个返回值，用“_”
	case "/":
		//return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		//panic("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 多个函数返回值
func div(a, b int) (q, r int) {
	/*q = a / b
	r = a % b
	return*/
	return a / b, a % b
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 重写返回int的pow方法
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
// 函数接收不定数量个参数，使用“...”
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(eval(3, 4, "/"))
	fmt.Println(eval(3, 4, "x"))
	q, r := div(13, 3)
	fmt.Println(q, r)

	if result, err := eval(3, 5, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, 3, 4))

	// 传入匿名函数
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
