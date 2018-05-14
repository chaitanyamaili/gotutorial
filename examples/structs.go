package main

import "fmt"

// defining struct Dog
type Dog struct {
	Bread  string
	Weight int
}

func main() {
	// assigning the values to Dog struct
	poodle := Dog{"Poodle", 32}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Bread: %v\nWeight: %v\n", poodle.Bread, poodle.Weight)
}
