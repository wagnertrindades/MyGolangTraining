package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type doubleZero struct {
	person
	licenseToKill bool
}

func main() {
	wagner := doubleZero{
		person: person{
			firstName: "Wagner",
			lastName:  "Trindade",
			age:       27,
		},
		licenseToKill: true,
	}

	fmt.Println(wagner.firstName, wagner.lastName, wagner.age, wagner.licenseToKill)
}
