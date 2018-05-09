package main

import (
	"fmt"
	"sort"
)

func main() {
	// slice of string type
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "White")
	fmt.Println(colors)

	// remove first element from the slice
	colors = append(colors[1:])
	fmt.Println(colors)

	// remove last element from the slice
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// slice of inter types
	numbers := make([]int, 5, 5)
	numbers[0] = 78
	numbers[1] = 13
	numbers[2] = 14
	numbers[3] = 12
	numbers[4] = 451
	fmt.Println(numbers)
	fmt.Println("Current numbers capacity", cap(numbers))

	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println("Current numbers capacity", cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)

	fmt.Println("Search in", sort.SearchInts(numbers, 100))
}
