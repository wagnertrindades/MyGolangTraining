package main

import "fmt"

func main() {
	var smallNumber int
	var largeNumber int

	fmt.Print("Please enter a small number: ")
	fmt.Scan(&smallNumber)

	fmt.Print("Please enter a large number: ")
	fmt.Scan(&largeNumber)

	remainder := largeNumber % smallNumber
	fmt.Println("The remainder of theses numbers is: ", remainder)
}
