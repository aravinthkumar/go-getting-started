package main

import "fmt"

// Constants can be declared outside the function similar to import statemetns
// iota can dynamically assign values to constants with iteration
// the value second is not assigned still iota is assigned at compile time.
const (
	first = iota
	second
	third = iota + 1
)

// New constant section with iota would be reset
const (
	fourth = iota
)

func main() {

	fmt.Println(first, second, third, fourth)

}
