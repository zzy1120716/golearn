// 分支流程控制语句：if、switch
package main

import (
	"fmt"
	"io/ioutil"
)

// switch后可以没有表达式
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d\n", score))
	case score < 60:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func main() {
	//const filename = "C:\\Users\\zzy\\go\\src\\golearn\\tutorial\\2-4\\abc.txt"
	const filename = "abc.txt"
	/*contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}*/
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// if 语句中赋值的变量作用域是当前if语句块，外部无法访问
	//fmt.Println(contents)		// error

	fmt.Println(eval(1, 2, "+"))
	//fmt.Println(eval(1, 2, "^"))

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		// 中断程序执行，其他分数也不输出了
		//grade(101),
		//grade(-3),
	)
}
