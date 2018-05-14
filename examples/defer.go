package main

import "fmt"

func main() {
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")

	myFunc()

	defer fmt.Println("Statement 1!")
	defer fmt.Println("Statement 2!")
	defer fmt.Println("Statement 3!")
	defer fmt.Println("Statement 4!")
	fmt.Println("Undefered Statement!")
}

func myFunc() {
	defer fmt.Println("Defered in the function!")
	fmt.Println("Not defered in the function!")
	myFunc2()
}

func myFunc2() {
	defer fmt.Println("Defered in the function2!")
	fmt.Println("Not defered in the function2!")
}
