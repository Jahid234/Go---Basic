package main

import "fmt"

func message(fname string, age int, dept string) {
	fmt.Println(fname, dept, age)
}

func main() {
	message("Jahid", 21, "CSE")
	message("Arafat", 21, "CSE")
	message("Pritom", 22, "BBA")
}
