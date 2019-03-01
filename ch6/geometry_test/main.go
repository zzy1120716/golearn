package main

import (
	"fmt"
	"golearn/ch6/geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q))
	//fmt.Printf("%T\n", Distance)
	fmt.Println(p.Distance(q))
	//fmt.Printf("%T\n", p.Distance)

	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
