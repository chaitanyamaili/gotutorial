package main

import "fmt"

func main() {
	var p *int

	if p != nil {
		fmt.Println("Value of p is", *p)
	} else {
		fmt.Println("p is nil")
	}

	var v int
	v = 42
	p = &v

	if p != nil {
		fmt.Println("Value of p is", *p)
	} else {
		fmt.Println("p is nil")
	}

	var v1 float64
	v1 = 56.56
	p1 := &v1
	fmt.Println("Value of p1 is", *p1)

	*p1 = *p1 / 11
	fmt.Println("Value of p1 is", *p1)
	fmt.Println("Value of v1 is", v1)
}
