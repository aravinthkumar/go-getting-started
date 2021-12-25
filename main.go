package main

import "fmt"

func main() {

	// Declare and initialize a map
	m := map[string]int{"foo": 42, "boo": 45, "coo": 48}
	fmt.Println(m)
	fmt.Println(m["foo"])
	// Change a element in map
	m["coo"] = 50
	fmt.Println(m)
	// Delete a element in map
	delete(m, "boo")
	fmt.Println(m)

}
