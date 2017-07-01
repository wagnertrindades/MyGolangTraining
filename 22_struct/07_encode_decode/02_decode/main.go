package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var wagner person
	reader := strings.NewReader(`{"First": "Wagner", "Last": "Trindade", "Age": 20}`)
	json.NewDecoder(reader).Decode(&wagner)

	fmt.Println(wagner.First)
	fmt.Println(wagner.Last)
	fmt.Println(wagner.Age)
}
