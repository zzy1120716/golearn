package main

import (
	"fmt"
	"golearn/ch2/tempconv"
)

/*var p = f()

func f() *int {
	v := 1
	return &v
}*/

func incr(p *int) int {
	*p++
	return *p
}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

// new可以重新定义为别的类型
func delta(old, new int) int { return new - old }

// 自动垃圾收集不由var或new声明变量的方式决定
var global *int

// 局部变量x从函数f中逃逸
func f() {
	var x int
	x = 1
	global = &x
}

// *y对于包来说是不可达的，可以马上被回收
func g() {
	y := new(int)
	*y = 1
}

func main() {

	/*x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)*/

	//fmt.Println(f() == f())

	//var x, y int
	//fmt.Println(&x == &x, &x == &y, &x == nil, *&x == *&y)

	//v := 1
	//incr(&v)	// 2
	//fmt.Println(incr(&v))	// 3

	/*p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)*/

	/*p := new(int)
	q := new(int)
	fmt.Println(p == q)	// "false"*/

	/*medals := []string{"gold", "silver", "bronze"}
	fmt.Println(medals)*/

	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZero)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.FToC(tempconv.CToF(tempconv.BoilingC)))
	fmt.Printf("%v\n", tempconv.CToK(tempconv.AbsoluteZero))
}
