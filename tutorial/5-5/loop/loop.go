// 循环语句
// 十进制转为二进制
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

// 传入参数类型由file变为io.Reader，适用范围更广了
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),          // 101
		convertToBin(13),         // 1101
		convertToBin(4534345234), // 100001110010001001001111000010010
		convertToBin(0),          // "", 省略特殊处理
	)

	printFile("C:\\Users\\zzy\\go\\src\\golearn\\tutorial\\2-4\\abc.txt")
	//printFile("abc.txt")
	s := `abc"d"
	kkkk
	123

	p`
	// 字符串当做文件来处理
	printFileContents(strings.NewReader(s))

	//forever()
}
