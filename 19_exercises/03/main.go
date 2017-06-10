package main

import "fmt"

func main() {
	greatest := greatestNumber(1, 5, 10, 28, 57, 1)
	fmt.Println(greatest)
}

func greatestNumber(numbers ...int) int {
	var largest int
	for _, number := range numbers {
		if number > largest {
			largest = number
		}
	}

	return largest
}
