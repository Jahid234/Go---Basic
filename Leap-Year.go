package main

import "fmt"

func main() {
	var year int
	fmt.Scanf("%v", &year)

	if year%400 == 0 || year%4 == 0 {
		fmt.Println("Leap year")
	} else {
		fmt.Println("Not Leap Year")
	}
}
