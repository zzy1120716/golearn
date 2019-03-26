// 验证欧拉公式
// 强制类型转换
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	// 求复数的模
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	// 让编译器不把i误认为变量，在前面加上"1"
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	euler()
	triangle()
}
