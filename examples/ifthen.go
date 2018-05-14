package main

import "fmt"

func main() {
	var x float64 = 42
	var result string

	// local variable for the condition
	if x := -2; x < 0 {
		result = "Less then zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println("Result:", result)
	fmt.Println("Value of x", x)
}
