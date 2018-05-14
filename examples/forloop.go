package main

import "fmt"

func main() {
	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	fmt.Println("\n")
	for i, v := range colors {
		//fmt.Println(colors[i])
		fmt.Println(i, ":", v)
	}

	fmt.Println("\n")
	sum = 1
	for sum < 1024 {
		sum += sum
		if sum == 256 {
			continue
		}
		if sum > 300 {
			goto endofcode
		}
		if sum > 500 {
			break
		}
		fmt.Println("Sum:", sum)
	}

	// jump to label
endofcode:
	fmt.Println("End of code")
}
