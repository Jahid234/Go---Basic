package main

import "fmt"

func main() {
	day := 8

	switch day {
	case 1, 8:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wednesday")
	case 6:
		fmt.Println("Thursday")
	case 7:
		fmt.Println("Friday")
	case 9:
		fmt.Println("Not Weekend")
	}

}
