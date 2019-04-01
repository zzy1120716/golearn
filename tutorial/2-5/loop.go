// 循环语句
// 十进制转为二进制
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	scanner := bufio.NewScanner(file)

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

	forever()
}
