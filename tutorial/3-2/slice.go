package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func printArray(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s := arr[2:6]
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	// 利用切片修改3-1中的printArray方法
	var arr1 [5]int
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println("printArray(arr1)")
	printArray(arr1[:])

	fmt.Println("printArray(arr3)")
	printArray(arr3[:])

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	// Slice的扩展，此处与python的切片不同
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	b := a[2:6]
	c := b[3:5]
	fmt.Println("b =", b) // [2 3 4 5]
	fmt.Println("c =", c) // [5 6]

	fmt.Println("a =", a)
	fmt.Printf("b=%v, len(b)=%d, cap(b)=%d\n",
		b, len(b), cap(b))
	// b=[2 3 4 5], len(b)=4, cap(b)=6
	fmt.Printf("c=%v, len(c)=%d, cap(c)=%d\n",
		c, len(c), cap(c))
	// c=[5 6], len(c)=2, cap(c)=3
	//fmt.Println(b[3:7])	// out of range
	fmt.Println(b[3:6]) // [5 6 7]
}
