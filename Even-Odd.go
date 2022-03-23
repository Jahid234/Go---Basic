package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%v", &n)

	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
