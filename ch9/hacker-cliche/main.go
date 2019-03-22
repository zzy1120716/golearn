package main

import "fmt"

func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

// GOMAXPROCS=1 go run golearn/ch9/hacker-cliche
// GOMAXPROCS=2 go run golearn/ch9/hacker-cliche
