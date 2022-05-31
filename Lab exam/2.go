/* Given two angles of a triangle by an user, calculate and third angle.
Make sure to check all possible error conditions if there is any   trianlge = 180 degree */

package main

import "fmt"

func main() {
	var a, b, triangle, answer float32

	fmt.Scanf("%f %f", &a, &b)

	triangle = 0
	answer = 180 - (a + b)

	if a == 0 || b == 0 {
		fmt.Println("Zero") /// Can be zero
	} else if a < 0 || b < 0 {
		fmt.Println("Negative") ///Can be negative
	} else if triangle >= answer {
		fmt.Println("Bigger") /// bigger
	} else {
		fmt.Println("Third angle  is: ", answer)
	}
}
