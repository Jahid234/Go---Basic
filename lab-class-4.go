package main

import "fmt"

func main() {

	// 	fmt.Println("Hello World")

	// 	//Lab class - 4

	// 	/*Static array, Dynamic array*/

	// 	var a1 = [3]int{1, 2, 3} // Static array
	// 	var a2 = []int{1, 2, 3}  // Dynamic array
	// 	a1[0] = 10

	// 	fmt.Println(a1)
	// 	fmt.Println(a2)

	// 	//Make function

	// 	var a3 = make([]int, 5, 10)
	// 	a3 = append(a3, 10)
	// 	fmt.Println(a3, len(a3), cap(a3))

	//MAP declaration

	var m1 = map[string]string{"un": "Jahid", "pwd": "12345"}

	fmt.Println(m1)

	var m2 = map[int]int{0: 0, 1: 10, 2: 20}

	fmt.Println(m2[0], m2[1], m2[2])

	//Another way

	var m3 = map[string]int{}
	m3["un"] = 1
	m3["pwd"] = 12345

	fmt.Println(m3["un"], m3["pwd"])

	//Make & Map

	var m4 = make(map[string]int, 3)
	m4["un"] = 2
	fmt.Println(m4)

	//Confusion 0(d) value

	var m5 = map[string]int{"a": 1, "b": 2}
	fmt.Println(m5)

	fmt.Println(m5["d"])

	//Comma.ok.Idiom

	var m6 = map[string]int{"a": 0, "b": 2}
	fmt.Println(m6)

	x, ok := m6["d"]
	x, ok = m6["a"]

	fmt.Println(x, ok)

}
