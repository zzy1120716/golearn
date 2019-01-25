package main

import "fmt"

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

	/*fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZero)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.FToC(tempconv.CToF(tempconv.BoilingC)))
	fmt.Printf("%v\n", tempconv.CToK(tempconv.AbsoluteZero))*/

	//p := new(int)
	//q := new(int)
	//fmt.Println(p == q)		// false

	//var x uint8 = 1<<1 | 1<<5	// 100010
	//var y uint8 = 1<<1 | 1<<2	// 110
	//fmt.Printf("%08b\n", x)	// 00100010，二进制8位，不足补0
	//fmt.Printf("%08b\n", y)	// 00000110
	//fmt.Printf("%08b\n", x&y)	// 00000010
	//fmt.Printf("%08b\n", x|y)	// 00100110
	//fmt.Printf("%08b\n", x^y)	// 00100100
	//fmt.Printf("%08b\n", x&^y)	// 00100000
	//
	//for i := uint(0); i < 8; i++ {
	//	if x&(1<<i) != 0 {
	//		fmt.Println(i)	// 1	5
	//	}
	//}
	//
	//fmt.Printf("%08b\n", x<<1)	// 01000100
	//fmt.Printf("%08b\n", x>>1)	// 00010001

	/*medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])	// "bronze", "silver", "gold"
	}*/

	//o := 0666
	//fmt.Printf("%d %[1]o %#[1]o\n", o)
	//x := int64(0xdeadbeef)
	//fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}
