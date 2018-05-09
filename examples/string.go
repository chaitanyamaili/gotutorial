package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly typed string"
	fmt.Printf("str1: %v: %T\n", str1, str1)

	str2 := "An explicitly typed string"
	fmt.Printf("str2: %v: %T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str2))

	lValue := "hello"
	uValue := "HELLO"
	fmt.Println("Equal?", (lValue == uValue))
	fmt.Println("Non case sensitive Equal?", strings.EqualFold(lValue, uValue))

	fmt.Println("Contains exp?", strings.ContainsAny(str1, "exp"))
	fmt.Println("Contains exp?", strings.Contains(str2, "exp"))

	fmt.Println("Count str1?", strings.Count(str1, "an"))
	fmt.Println("Count str2?", strings.Count(str2, "exp"))

	fmt.Println("Compare?", strings.Compare("7", "5"))

	s := []string{"foo", "bar", "baz"}
	joinS := strings.Join(s, ", ")
	fmt.Println(joinS)
	fmt.Println(strings.Split(joinS, ", "))
}
