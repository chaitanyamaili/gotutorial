package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.txt")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	myError := errors.New("My error string")
	fmt.Println(myError.Error())

	attendance := map[string]bool{
		"Ann":  true,
		"Mike": false}

	attended, ok := attendance["Ram"]

	if ok {
		fmt.Println("Ram attended?", attended)
	} else {
		fmt.Println("No record found for Ram")
	}
}
