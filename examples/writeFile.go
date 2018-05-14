package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go!"

	file, err := os.Create("./examples/fromString.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with file of %v characters\n", ln)

	bytes := []byte(content)
	ioutil.WriteFile("./examples/fromBytes.txt", bytes, 0644)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
