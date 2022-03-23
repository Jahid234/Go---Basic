package main

import (
	"fmt"
)

const b int = 9 // typed constant

const c = 10 //untyped constant

const (
	A int = 1
	B     = 3
	C     = 10.00
)

func main() {
	const a int = 1

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
