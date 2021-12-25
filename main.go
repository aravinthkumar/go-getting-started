package main

import "fmt"

func main() {
	// Simple example of constant, constant can't be changed hence Go expects to declare the value along.
	const pi = 3.1415
	fmt.Println(pi)

	// Dynamic changing of datatype
	const c = 3
	fmt.Println(c + 3)

	fmt.Println(c + 1.2)

	// Strict type on constant
	const d int = 4
	fmt.Println(d + 3)
	// casting to float when dealing with float value else go doesn't allow
	fmt.Println(float32(d) + 1.2)
}
