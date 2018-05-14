package main

import (
	"fmt"

	"github.com/chaitanyamaili/gotutorial/stringutil"
)

func main() {
	n, l := stringutil.FullName("Chaitanya", "Maili")
	fmt.Printf("Full: %v, length: %v\n", n, l)

	n1, l1 := stringutil.FullNameNakedReturn("Ram", "Kumar")
	fmt.Printf("Full: %v, length: %v\n", n1, l1)
}
