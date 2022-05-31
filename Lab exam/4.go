/*Write a program that sums all numbers from 1 to 30 which are divisible by either 3 or 5
 */

package main

import "fmt"

func main() {
	var sum int
	sum = 0

	for i := 1; i <= 30; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}

	fmt.Println(sum)
}
