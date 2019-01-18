package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)	// each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler 响应请求URL r 的 Path 组件
func handler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}