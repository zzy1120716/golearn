package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex

var i = 0
var max = 300

func Print(content string, divNum int, wg *sync.WaitGroup) {
	for i <= max {
		m.Lock()
		if i%3 == divNum {
			fmt.Print(content)
			i++
		}
		m.Unlock()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go Print("a", 0, &wg)
	go Print("l", 1, &wg)
	go Print("i", 2, &wg)

	wg.Wait()
}
