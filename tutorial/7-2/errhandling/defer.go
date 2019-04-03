package main

import (
	"bufio"
	"errors"
	"fmt"
	"golearn/tutorial/7-1/fib"
	"os"
)

func writeFile(filename string) {
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("this is a custom error")
	if err != nil {
		//fmt.Println("file already exists")
		//fmt.Println("Error:", err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}
