package main

import "fmt"

func main() {

	half := func(number int) (int, bool) {
		dividedBy2 := number / 2

		var even bool
		if number%2 == 0 {
			even = true
		}

		return dividedBy2, even
	}

	fmt.Println(half(1))
	fmt.Println(half(2))
}
