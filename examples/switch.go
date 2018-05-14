package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(6) + 1
	// fmt.Println("Day", dow)

	result := ""

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "It's Sunday!"
		//result = strings.Join(["It's Sunday!",strconv.FormatInt(dow)], ": ")
	case 7:
		result = "It's Saturday!"
	default:
		result = "It's Weekday!"
		//result = strings.Join({"It's weekday!",strconv.FormatInt(dow)}, ": ")
	}
	fmt.Println(result)

	x := 42
	switch {
	case x > 0:
		result = "Greater than zero"
		// fallthrough
	case x == 0:
		result = "Equal to zero"
	case x < 0:
		result = "Less than zero"
	}
	fmt.Println(result)
}
