package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price string
}

func main() {
	url := "http://www.services.explorecalifornia.org/json/tours.php"

	content := contentFromServer(url)
	//fmt.Println(content)
	tours := toursFromJson(content)
	//fmt.Printf("Type: %T\n", tours)
	//fmt.Printf("%+v\n", tours)
	//fmt.Println(tours)

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v {$%.2f}\n", tour.Name, price)
	}
}

func contentFromServer(url string) string {
	resp, err := http.Get(url)
	checkError(err)

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}
