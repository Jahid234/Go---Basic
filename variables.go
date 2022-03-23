package main

import "fmt"

func main() {

	// variable declaration with initial value

	var student1 string = "John"
	var student2 = "Doe" //infarred
	x := 2               //infarred
	y := 2.04            // infarred

	var student3 string = "Jahid"

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(student3)

	// Variable declaration without initial value

	var a string
	var b int
	var c float32
	var d bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
