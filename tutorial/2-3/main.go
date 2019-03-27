// 常量和枚举
package main

import (
	"fmt"
	"math"
)

// 不要使用大写，类型自动推断
//const filename = "abc.txt"
const (
	filename = "abc.txt"
	a, b     = 3, 4
)

func consts() {
	//const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	// iota自增值种子，利用表达式，进行自增值枚举
	const (
		/*cpp = 0
		java = 1
		python = 2
		golang = 3*/
		cpp = iota
		_
		java
		python
		golang
		javascript
	)
	fmt.Println(cpp, java, python, golang, javascript)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()
	enums()
}
