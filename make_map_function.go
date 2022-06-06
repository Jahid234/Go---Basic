package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// - make()
	var m1 = make([]int, 0, 2)

	m1 = append(m1, 2)
	m1 = append(m1, 2)
	m1 = append(m1, 2)
	m1 = append(m1, 3)
	m1 = append(m1, 3)
	m1 = append(m1, 3)
	m1 = append(m1, 3)

	fmt.Println(m1, cap(m1), len(m1), m1[1])

	// - map[]
	m2 := map[string]int{
		"Jahid":  200103020029,
		"Pritom": 200103020063,
	}

	m3 := map[int]string{
		200103020062: "Jamil",
		200103020013: "Arafat",
	}

	fmt.Println(m2["Jahid"])

	fmt.Println(m3[200103020062])

	// use - map() & make[]

	m4 := make(map[string]int)

	m4["Jahid"] = 200103020029
	m4["Jamil"] = 200103020062

	fmt.Println(m4["Jahid"])

	// -empty map(using make() and without make())
	m5 := make(map[string]int)
	var m6 map[string]int

	fmt.Println(m5 == nil)
	fmt.Println(m6 == nil)

	// -In map adding an elements & updating an elements

	m4["Jahid"] = 2001030200290
	m4["Pritom"] = 200103020063

	fmt.Println(m4)

	// -In map delete an element

	delete(m4, "Jahid")

	fmt.Println(m4)

	// -Existing key & value check

	_, ok1 := m2["Jahid"]
	val1, ok2 := m2["Pritom"]

	fmt.Println(ok1, val1, ok2)

}
