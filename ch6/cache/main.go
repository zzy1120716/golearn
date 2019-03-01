package main

import (
	"fmt"
	"sync"
)

/*var (
	mu sync.Mutex // guards mapping
	mapping = make(map[string]string)
)
func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}*/

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	cache.mapping["a"] = "aaa"
	fmt.Println(Lookup("a"))
}
