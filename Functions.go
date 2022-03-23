package main

import "fmt"

func Jahid() {
	fmt.Println("I am Jahid")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}

func main() {
	Jahid()
}
