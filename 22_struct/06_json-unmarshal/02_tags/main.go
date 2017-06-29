package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int `json:"Wisdom score"`
}

func main() {
	var wagner person
	byteSlice := []byte(`{"First": "Wagner", "Last": "Trindade", "Wisdom score": 27}`)
	json.Unmarshal(byteSlice, &wagner)

	fmt.Println(wagner.First)
	fmt.Println(wagner.Last)
	fmt.Println(wagner.Age)
}
