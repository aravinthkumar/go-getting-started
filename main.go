package main

import "fmt"

func main() {
	// Holding the address(pointer) and memory
	var firstName *string = new(string)
	// deferencing the firstName to add value
	*firstName = "Aravinth"
	// Printing the deference value and address
	fmt.Println(*firstName, firstName)

	// Another example with address of operator(pointer)
	lastName := "Kumar"
	ptr := &lastName
	fmt.Println(lastName, ptr)

	// Trying to change the value of lastName but the pointer values remains the same
	lastName = "Jain"
	fmt.Println(lastName, ptr)
}
