package main

import "fmt"

func main() {

	arr := [3]int{4, 5, 5}
	fmt.Println(arr, arr[1])

	// Slice is initialized and declared from a array
	slice := arr[:]
	fmt.Println(slice)

	// Array and Slice is changed the data is same since they point to the same value from the memory
	arr[1] = 9
	slice[2] = 10

	fmt.Println(arr, slice)

	// Slice declared and initialzed without an array
	slice1 := []int{8, 9, 0}
	fmt.Println(slice1)
	// Slice is appended with more element
	slice1 = append(slice1, 4, 3)
	fmt.Println(slice1)

	// Slice options on a collections
	slice2 := slice1[:1]
	slice3 := slice1[1:2]
	slice4 := slice1[1:]
	fmt.Println(slice1, slice2, slice3, slice4)
}
