package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	first         string
	licenseToKill bool
}

func (p person) greetings() string {
	return "Hello, I'm a normal man."
}

func (d doubleZero) greetings() string {
	return "Hello, I'm a super man."
}

func main() {
	// The field in struct override inner field in other struct
	james := doubleZero{
		person:        person{"James", "Bond", 27},
		first:         "Zé",
		licenseToKill: true,
	}
	fmt.Println(james.first, james.last, james.age, james.licenseToKill)

	// The method in struct override inner method in other struct
	fmt.Println(james.greetings())
	fmt.Println(james.person.greetings())
}
