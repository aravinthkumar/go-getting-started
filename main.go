package main

import (
	"fmt"
	"github/aravinthkumar/go-getting-started/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Aravinth",
		LastName:  "Kumar",
	}
	fmt.Println(u)
}
