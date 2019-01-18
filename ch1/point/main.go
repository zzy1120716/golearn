package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	var p1 = Point{1, 1}
	var p2 = Point{2, 2}
	fmt.Println(p1)
	fmt.Println(p2)
}