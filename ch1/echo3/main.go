package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(strings.Join(os.Args[:], " "))
	fmt.Println(os.Args[1:])
	for i, arg := range os.Args[1:] {
		fmt.Println(arg + " " + strconv.Itoa(i))
	}
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d %s", i, arg + "\n")
	}
}
