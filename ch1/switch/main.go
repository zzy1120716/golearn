package main

import (
	"fmt"
	"math/rand"
)

func main() {
	heads := 0
	tails := 0
	for i := 0; i < 10000; i++ {
		switch coinflip() {
		case "heads":
			heads++
		case "tails":
			tails++
		default:
			fmt.Println("landed on edge!")
		}
	}
	fmt.Printf("正面次数为：%d", heads)
	fmt.Println()
	fmt.Printf("反面次数为：%d", tails)
	fmt.Println()
	fmt.Printf("站立次数为：%d", 10000 - heads - tails)
	fmt.Println()
	fmt.Printf("%d", Signum(1))
	fmt.Println()
	fmt.Printf("%d", Signum(0))
	fmt.Println()
	fmt.Printf("%d", Signum(-1))
}

func coinflip() string {
	var result = [2]string{"heads", "tails"}
	return result[rand.Intn(2)]
}

// tagless switch
func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}