package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 23, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Go launced at %s\n", t)

	now := time.Now()
	fmt.Printf("Current time is %s\n", now)

	fmt.Println("The month is", now.Month())
	fmt.Println("The day is", now.Day())
	fmt.Println("The weekday is", now.Weekday())
	fmt.Println("The unix timestamp is", now.Unix())

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))
	shortFormat := "2/1/2006"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
}
