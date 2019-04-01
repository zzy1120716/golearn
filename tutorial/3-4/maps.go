package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok) // golang true

	// key not in map
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName, ok)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok) // ccmouse true

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok) // "" false
}
