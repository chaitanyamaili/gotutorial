package main

import "fmt"

func main() {
	fooFunc()
	fmt.Println(addFunc(10, 23))

	sum := addAll(12, 54, 59)
	fmt.Println(sum)
}

func fooFunc() {
	fmt.Println("Inside the foo")
}

// if both the parameters are of same type
func addFunc(v1, v2 int) int {
	return v1 + v2
}

func addAll(values ...int) int {
	sum := 0
	for k := range values {
		sum += values[k]
	}
	fmt.Printf("%T\n", values)
	return sum
}
