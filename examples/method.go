package main

import "fmt"

type Dog struct {
	Bread  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d *Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println("\n", d.Sound)
}

func main() {
	rocky := Dog{"Poodle", 29, "Woof Woof"}
	fmt.Println(rocky)
	rocky.Speak()
	rocky.SpeakThreeTimes()

	rocky.Sound = "Ark"
	rocky.Speak()
	rocky.SpeakThreeTimes()
	rocky.SpeakThreeTimes()
	rocky.SpeakThreeTimes()
}
