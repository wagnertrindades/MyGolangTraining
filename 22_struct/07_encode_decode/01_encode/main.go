package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	wagner := person{"Wagner", "Trindade", 27, 007}
	json.NewEncoder(os.Stdout).Encode(wagner)
}
