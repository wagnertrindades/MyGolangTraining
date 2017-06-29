package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	wagner := Person{"Wagner", "Trindade", 27}
	fmt.Println(wagner.first, wagner.last, wagner.age)
	byteSlice, _ := json.Marshal(wagner)
	fmt.Println(byteSlice)
	fmt.Println(string(byteSlice))
}
