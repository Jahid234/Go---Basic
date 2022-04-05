package main

import "fmt"

func main() {

	/*...........Lab - class - 5.............*/

	//GO - Structure

	type student struct {
		firstname  string
		familyname string
		id         int
	}

	//x := student{"Jahidul", "Islam", 1}

	x := student{
		familyname: "Islam", firstname: "Jahid", id: 1}

	fmt.Println(x)

	/*......Go - Control - Structure.......*/

	//Simple If

	A := 10

	if A > 3 {
		fmt.Println("Hello")
	} else {
		fmt.Println("World")
	}

	//Nested if

	if A%2 == 0 {
		if A%5 == 0 {
			fmt.Println("Divisible by 2 and 5.")
		}
		fmt.Println("Divisible by 2, but not 5")
	} else {
		fmt.Println("Not divisible by 2")
	}

	//ladder(if, else if, else)

	B := 20
	if B%2 == 0 && B%5 == 0 {
		fmt.Println("Divsisible by 2 nad 5")
	} else if B%2 == 0 && B%5 != 0 {
		fmt.Println("Divisible by 2, but not 5")
	} else {
		fmt.Println("Not divisible by 2")
	}

	//for - loop

	for i := 0; i < 10; i++ {
		fmt.Println("I want to join GOOGLE.")
	}

	//while - loop

	var i int
	for i < 10 {
		fmt.Println("I want to join GOOGLE.")
		i++
	}

	// for - range loop ....(if (value or index) not needed use underscore)

	var X = []int{5, 10, 15, 20, 25, 30}

	for i, v := range X {
		fmt.Println(i, v)
	}

	//use underscore

	var C = []int{5, 20, 30}

	for _, v := range C {
		fmt.Println(v)
	}
}
