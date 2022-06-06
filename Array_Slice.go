package main

import "fmt"

func main() {

	// -Array
	arr := [...]int{2, 3, 4}
	arr1 := [...]string{"Jahid", "Pritom", "Jamil"}

	// -change the values
	arr[2] = 5

	fmt.Println(arr)
	fmt.Println(arr1)

	///	-specific elements initialize

	arr3 := [...]int{1: 2, 2: 2}
	fmt.Println(arr3)

	/// - Slice
	s1 := []int{2, 3, 4, 5}
	fmt.Println(s1, len(s1), cap(s1))

	/// - Create a slice from an array

	arr2 := [...]int{5, 6, 7, 8}
	s2 := arr2[2:4]

	fmt.Println(s2)

	/// -In slice append

	s1 = append(s1, 7, 8, 9)
	fmt.Println(s1)

	/// -Create a slice with make()

	s3 := make([]int, 0, 3)
	s3 = append(s3, 1)
	s3 = append(s3, 2)
	s3 = append(s3, 3)
	fmt.Println(s3)

	/// -Append one slice to another slice

	s4 := []int{10, 11, 12}

	s1 = append(s1, s4...)

	fmt.Println(s1)

	// -Change the length of a slice
	s2 = arr2[1:3]
	fmt.Println(s2)
}
