package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{"Wagner", "Trindade", 27}
	fmt.Println(p1.firstName, p1.lastName, p1.age)
}
