package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%v", &n)

	c := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			c++
		}
	}

	if c == 2 {
		fmt.Println("Prime")
	} else {
		fmt.Println("Non-Prime")
	}

}
