package main

import "fmt"

func main() {
	var a, b, c, d int = 2, 4, 5, 6

	// When type keyword not use then variable declaration in different line no problem

	var x, y = 7, "hello"
	t, s := 9, "World!"

	fmt.Println(a, b, c, d)
	fmt.Println(x, y, t, s)

	// Variable declaration in a block

	var (
		f int
		g int = 10

		h = "Arisha"
	)

	fmt.Println(f, g, h)
}
