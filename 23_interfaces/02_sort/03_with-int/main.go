package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{1, 49, 98, 98, 2, 27, 3, 29, 4, 938, 0}
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)

	// Other way to sort slice of ints
	i := []int{1, 49, 98, 98, 2, 27, 3, 29, 4, 938, 0}
	sort.Ints(i)
	fmt.Println(i)
}
