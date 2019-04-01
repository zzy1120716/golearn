package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray1(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	// 不指定数量，“...”不可省略
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 数组的遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	// 传入值
	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)
	//printArray(arr2)	// error

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)

	// 传入指针
	fmt.Println("printArray(arr1)")
	printArray1(&arr1)

	fmt.Println("printArray(arr3)")
	printArray1(&arr3)
	//printArray(arr2)	// error

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
