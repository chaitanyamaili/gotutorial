package main

import "fmt"

func main() {
	var colors [4]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	colors[3] = "Orange"
	fmt.Println(colors)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println("Count in colors", len(colors))
	fmt.Println("Count in numbers", len(numbers))
}
