/* Write a program that takes minutes as input and print the total number of hours
and reamianing minutes.*/

package main

import "fmt"

func main() {
	var min, hour, rem_min int
	fmt.Scanf("%d", &min)

	hour = min / 60
	rem_min = min % ((min / 60) * 60)

	fmt.Println(hour, rem_min)
}
