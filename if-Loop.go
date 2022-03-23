package main

import "fmt"

func main() {
	var x int
	fmt.Scanf("%x", &x)

	if x%5 == 0 {
		fmt.Println("Divisible")
	} else {
		fmt.Println("Not Divisible")
	}
}
