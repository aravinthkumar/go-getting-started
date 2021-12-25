package main

import "fmt"

func main() {
	// Definition of a struct
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	// Initialize a variable with the struct type, following would have initial values
	var user1 user
	fmt.Println(user1)

	// Initialize with values
	user1.ID = 1
	user1.FirstName = "Aravinth"
	user1.LastName = "Kumar"

	fmt.Println(user1)

	fmt.Println(user1.FirstName, user1.LastName)
	// Another way to Declare and initialize a variable with values
	user2 := user{
		ID:        2,
		FirstName: "Rahul",
		LastName:  "Jain",
	}
	fmt.Println(user2)

}
