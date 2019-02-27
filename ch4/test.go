package main

import (
	"fmt"
	"time"
)

func main() {
	/*var a [3]int		// array of 3 integers
	fmt.Println(a[0])
	fmt.Println(a[len(a) -  1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}*/

	/*var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[2])
	fmt.Println(r[2])*/

	/*q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)

	q = [4]int{1, 2, 3, 4}	// wrong*/

	/*type Currency int

	const (
		USD Currency = iota	// 美元
		EUR					// 欧元
		GBP					// 英镑
		RMB					// 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	fmt.Println(RMB, symbol[RMB])	// "3 ¥"*/

	/*r := [...]int{99: -1}
	for i, v := range r {
		fmt.Printf("%d %d\n", i, v)
	}*/

	/*a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)		// "true false false"
	d := [3]int{1, 2}*/
	//fmt.Println(a == d)		// error

	/*months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months[0])		// ""
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)			// [April May June]
	fmt.Println(summer)		// [June July August]

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appers in both\n", s)
			}
		}
	}

	// fmt.Println(summer[:20])	// panic: runtime error: slice bounds out of range
	endlessSummer := summer[:5]		// extend a slice (within capacity)
	fmt.Println(endlessSummer)		// [June July August September October]*/

	/*var s []int
	fmt.Printf("%t %t\n", len(s) == 0, s == nil)
	s = nil
	fmt.Printf("%t %t\n", len(s) == 0, s == nil)
	s = []int(nil)
	fmt.Printf("%t %t\n", len(s) == 0, s == nil)
	s = []int{}
	fmt.Printf("%t %t\n", len(s) == 0, s == nil)

	s = make([]int, 3)
	s = make([]int, 3, 5)
	s = make([]int, 5)[:3]*/

	/*var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)	// ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']*/

	/*ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}*/
	//ages := make(map[string]int)	// mapping from strings to ints
	/*ages := map[string]int{}
	ages["alice"] = 31
	ages["charlie"] = 34
	ages["alice"] = 32
	fmt.Println(ages["alice"])
	delete(ages, "alice")
	fmt.Println(ages)
	ages["bob"] = ages["bob"] + 1
	ages["bob"] += 1
	ages["bob"]++
	fmt.Println(ages)
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}*/

	/*ages := map[string]int{"alice": 31, "bob": 20, "charlie": 34}
	names := make([]string, 0, len(ages))
	fmt.Println(names)*/

	/*var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)*/

	//ages["carol"] = 21		// panic: assignment to entry in nil map

	/*age, ok := ages["bob"]
	if !ok {
		fmt.Println("\"bob\" is not a key in this map: age == 0.")
		fmt.Println(age)
	}
	if age, ok := ages["bob"]; !ok {
		fmt.Println("\"bob\" is not a key in this map: age == 0.")
		fmt.Println(age)
	}*/

	/*x := map[string]int{"alice": 31, "bob": 20, "charlie": 34}
	y := map[string]int{"alice": 31, "bob": 20, "charlie": 34}
	fmt.Println(equal(x, y))*/

	/*var dilbert Employee
	fmt.Printf("%T", dilbert)

	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	//(*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason*/

	fmt.Println(Scale(Point{1, 2}, 5)) // {5 10}

	//pp := &Point{1, 2}
	/*	pp := new(Point)
	*pp = Point{1, 2}*/
	/*fmt.Println(pp)

	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q) // "false"*/

	/*type address struct {
		hostname string
		port int
	}

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)*/

	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Println(w)
}

type Point struct{ X, Y int }

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

func EmployeeByID(id int) *Employee { return nil }

var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func zero1(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

// slice比较相等
/*func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}*/

// map比较相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
