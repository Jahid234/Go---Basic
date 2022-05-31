/* Write a program that takes two numbers from the user and then print
the square of their summation. */

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	sum := a + b
	fmt.Println("Summation is: ", sum*sum)
}
