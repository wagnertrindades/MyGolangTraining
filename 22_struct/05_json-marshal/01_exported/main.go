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
	wagner := Person{"Wagner", "Trindade", 27}
	byteSlice, _ := json.Marshal(wagner)
	fmt.Println(byteSlice)
	fmt.Printf("%T \n", byteSlice)
	fmt.Println(string(byteSlice))
}
