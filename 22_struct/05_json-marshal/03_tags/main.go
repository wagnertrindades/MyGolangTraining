package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"Wisdom score"`
}

func main() {
	wagner := Person{"Wagner", "Trindade", 27}
	byteSlice, _ := json.Marshal(wagner)
	fmt.Println(string(byteSlice))
}
