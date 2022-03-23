package main

import "fmt"

func fact(n int) int {
	var fact int = 1

	fmt.Scanf("%v", &n)

	for i := 1; i <= n; i++ {
		fact = fact * i

	}
	return fact
}

func main() {
	result := fact(5)
	fmt.Println("Factorial: ", result)
}
