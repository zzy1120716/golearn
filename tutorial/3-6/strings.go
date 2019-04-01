package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!" // UTF-8
	fmt.Println(len(s))
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	//fmt.Printf("%X\n", []byte(s))
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// 字符串中有多少个rune
	fmt.Println(utf8.RuneCountInString(s))

	// 依次打印每个字符（支持多语言）的方法
	bytes := []byte(s)
	for len(bytes) > 0 {
		// 英文size为1，中文size为3
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 按下标取字符（直接转为rune数组）
	// 底层新开辟一个rune数组，再进行处理
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

	// strings包，可实现字符串基本操作
	// Fields, Split, Join
	// Contains, Index
	// ToLower, ToUpper
	// Trim, TrimRight, TrimLeft

	// 分割字符串，多个空格
	fmt.Printf("%q\n", strings.Fields("Hello  world ! My name  is J."))
	fmt.Printf("%q\n", strings.Split("Hello  world ! My name  is J.", " "))
}
