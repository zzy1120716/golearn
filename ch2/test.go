package main

import "fmt"

var p = f()

func f() *int {
	v := 1
	return &v
}

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

	p := new(int)
	q := new(int)
	fmt.Println(p == q)	// "false"
}
