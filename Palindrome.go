package main

import "fmt"

func main() {
	var n, n1, rem int
	fmt.Scanf("%n", &n)
	n1 = n

	rev := 0

	for n > 0 {
		rem = n % 10
		rev = rev*10 + rem
		n = n / 10
	}

	if n1 == rev {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
