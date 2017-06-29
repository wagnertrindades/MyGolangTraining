package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var wagner Person
	fmt.Println(wagner.First)
	fmt.Println(wagner.Last)
	fmt.Println(wagner.Age)

	fmt.Println("-----------------")

	byteSlice := []byte(`{"First": "Wagner", "Last": "Trindade", "Age": 27}`)
	json.Unmarshal(byteSlice, &wagner)

	fmt.Println(wagner.First)
	fmt.Println(wagner.Last)
	fmt.Println(wagner.Age)
}
