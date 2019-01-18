package main

import (
	"bufio"
	"fmt"
	"golearn/ch2/unitconv"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			num, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			conv(num)
		}
	} else {
		for _, arg := range os.Args[1:] {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			conv(num)
		}
	}
}

func conv(num float64) {
	m := unitconv.Meter(num)
	ft := unitconv.Foot(num)
	kg := unitconv.Kilogram(num)
	lb := unitconv.Pound(num)
	fmt.Printf("length: %s = %s, %s = %s\n", m, unitconv.MToFt(m), ft, unitconv.FtToM(ft))
	fmt.Printf("weight: %s = %s, %s = %s\n", kg, unitconv.KgToLb(kg), lb, unitconv.LbToKg(lb))
}
