package main

import "fmt"

func main() {
	fmt.Println(greet("Ai", "Ou"))
}

func greet(firstName, lastName string) (string, string) {
	return fmt.Sprint(firstName), fmt.Sprint(lastName)
}
