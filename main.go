package main

import "fmt"

func main() {

	// Normal declaration and assignment
	var i int
	i = 4
	fmt.Println(i)

	// Normal declaration with assignment
	var f float32 = 3.14
	fmt.Println(f)

	// Go determines the variable type during assignment
	firstName := "Arman"
	fmt.Println(firstName)

	// defining a boolean
	b := true
	fmt.Println(b)

	// Go built in complex function
	c := complex(3, 4)
	fmt.Println(c)

	// to print the real and imaginary and also decalare two variable in single statement
	r, im := real(c), imag(c)
	fmt.Println(r, im)

	//Another example for declaring two variable in same statement.
	middleName, lastName := "Malik", "Khan"
	fmt.Println(middleName, lastName)

}
