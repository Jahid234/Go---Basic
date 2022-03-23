package main

import "fmt"

func main() {

	/*fmt.Println("Hello World")

	//Variables

	var x int
	x = 10
	y := 20
	fmt.Println("Data: ", x, y)

	z := x + y
	s := y - x
	t := y * x
	d := y / x

	fmt.Println("Sum: ", z, "SUbtract: ", s, "Multiply: ", t, "Divide: ", d)*/

	//Datatypes(int, float, Rune & String, Complex, Boolean)
	//Compile time vs run time
	// compile time = compile time error, runtime = in run time error
	//Arithmetic operator, relational operator, bitwise operator

	/* Arithmetic -> +, -, *, /, %, ++, --
	   Relational -> ==, !=, <=, >=, <, >
	   Logical -> &&, ||, !
	   Bitwise -> &, |, ^, <<, >>
	   Assignment -> =, +=, -=, /=, *=, &=, |=, ^=, <<=, >>=
	*/

	var A int = 10
	var B int = 20

	fmt.Println(A == B)
	fmt.Println(A << 1)

	fmt.Println(A)

	if true {
		fmt.Println("Hello")
		fmt.Println(A << 1) // right shift
		fmt.Println(A >> 1) // left shift
		fmt.Println(A & B)  // Bitwise and
		fmt.Println(A | B)  //Bitwise or
		fmt.Println(A ^ B)  //Bitwise xor
	}

	// Task- 1 take a number print whether it is divisible by 3

	/*var a int
	fmt.Scanf("%a", &a)

	if a%3 == 0 {
		fmt.Println("Divisible")

	} else {
		fmt.Println("Not divisible")
	}*/

	// datatypes float
	/*var a float64 = 10.5
	b := 10
	var p float64 = 10.5
	var c int = int(a) + b // typecasting
	fmt.Println(c)
	fmt.Println(a == p)*/

	///  10.5
	/// 105/100
	/// not possible in float64
	// var z float64 = 105 / float64(10) , not possible
}
