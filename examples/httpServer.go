package main

import (
	"fmt"
	"net/http"
)

type HelloStruct struct{}

func (h HelloStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello from Go web server!!!<h1>")
}

func main() {
	var h HelloStruct
	err := http.ListenAndServe("localhost:4000", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
