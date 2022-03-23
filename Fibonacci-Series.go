package main

import "fmt"

func FibonacciSeries(num int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("Series is: %d %d", a, b)

	for true {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)

	}
}

func main() {
	FibonacciSeries(10)
	FibonacciSeries(16)
	FibonacciSeries(100)
}
