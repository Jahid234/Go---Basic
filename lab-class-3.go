package main

import "fmt"

func main() {
	//Lab-class-3

	// //Decalare an array

	// var a [3]int

	// var x = [3]int{1, 2, 3} ///Array literal
	// x = [3]int{5, 6, 7}

	// var y = [...]int{6, 7, 8}

	// //Multi dimensional array declaration
	// var arr [3][4]int

	// arr[0] = [...]int{1, 2, 3, 4}
	// fmt.Println(arr)

	// fmt.Println(x)
	// fmt.Println(a)
	// fmt.Println(y)

	// fmt.Println("Lab-class-3")

	// //Declare a slice(no dot declare in size)

	// var B = []int{1, 2, 3}
	// B = append(B, 4)
	// B = append(B, 5)
	// fmt.Println(B)

	// //Make function

	// //var X = make([]int, 0, 10)
	// var X = make([]int, 0, 2)

	// fmt.Println(X, len(X), cap(X))
	// X = append(X, 1)
	// fmt.Println(X, len(X), cap(X))

	/// very important Slicing

	var x = make([]int, 0, 4)
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	x = append(x, 4)
	fmt.Println(x)

	var y = make([]int, 2)

	//copy(y, x)
	// x[startindex : endindex+1]

	copy(y, x[1:3])
	fmt.Println(y)

}
