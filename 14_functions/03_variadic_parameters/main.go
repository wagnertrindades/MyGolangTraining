package main

import "fmt"

func main() {
	n := averange(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func averange(sliceFunc ...float64) float64 {
	fmt.Println(sliceFunc)
	fmt.Printf("%T", sliceFunc)

	var total float64
	for _, value := range sliceFunc {
		total += value
	}

	return total / float64(len(sliceFunc))
}
