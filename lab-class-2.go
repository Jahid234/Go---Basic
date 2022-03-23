package main

import "fmt"

func main() {
	///Rune data types
	//variables
	//Multiple variavbles

	// var x rune
	// x = 98

	// var y rune
	// var z rune
	// z = '\U000009b8'
	// fmt.Printf("%c %c", x, y)
	// fmt.Printf("%c", z)

	//shadow variables

	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		//Shadowing
		fmt.Println(x)

	}
	//Shadowed
	fmt.Println(x)

	//Single constants

	const b int = 9 // typed constant

	const c = 10 //untyped constant

	// Multiple Constants

	const (
		A int = 1
		B     = 3
		C     = 10.00
	)

	const a int = 1

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
