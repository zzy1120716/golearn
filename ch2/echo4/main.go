// 打印命令行参数
// 使用flag包，解析命令行参数，获取标志参数
// 在输入非法参数后，给出相应提示
// go run main.go a bc def
// go run main.go -s / a bc def
// go run main.go -n a bc def
// go run main.go -help
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	// 先调用flag.Parse，更新每个标志参数对应变量的值
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}