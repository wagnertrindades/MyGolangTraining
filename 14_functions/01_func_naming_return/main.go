package main

import "fmt"

func main() {
	fmt.Println(greet("Ai", "Ou"))
}

func greet(firstName, lastName string) (names string) {
	names = fmt.Sprint(firstName, lastName)
	return
}
