package main

import "fmt"

func main() {
	n, l := fullName("Chaitanya", "Maili")
	fmt.Printf("Full: %v, length: %v\n", n, l)

	n1, l1 := fullNameNakedReturn("Ram", "Kumar")
	fmt.Printf("Full: %v, length: %v\n", n1, l1)
}

func fullName(f, l string) (string, int) {
	full := f + " " + l
	len := len(full)
	return full, len
}

func fullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
