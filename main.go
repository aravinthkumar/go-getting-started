package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers")
	port := 3000
	// function invoked with argument
	b, err := startWebServer(port, "2")
	fmt.Println(b, err)
}

// Function with input parametes
func startWebServer(port int, numberOfRetries string) (bool, error) {

	fmt.Println("Starting server ....")

	fmt.Println("Number of time retried", numberOfRetries)
	// do something
	fmt.Println("Server started in port", port)

	return true, nil

}
