package main

import "fmt"

func main() {
	// Array declaration of type int
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr[2], arr)

	// Another way to declare and initialize
	arr2 := [3]int{4, 5, 5}
	fmt.Println(arr2, arr2[1])

}
