package main

import "fmt"

func main() {
	fmt.Println(half(6))
	fmt.Println(half(3))
}

func half(number int) (int, bool) {
	dividedBy2 := number / 2

	var even bool
	if number%2 == 0 {
		even = true
	}

	return dividedBy2, even
}
