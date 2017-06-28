package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	wagner := person{"Wagner", "Trindade", 27}
	james := person{"James", "Bond", 37}
	fmt.Println(wagner.fullName())
	fmt.Println(james.fullName())
}
