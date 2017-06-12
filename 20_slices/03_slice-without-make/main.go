package main

import "fmt"

func main() {

	mySlice := []int{1, 4, 19, 20}

	fmt.Println("----------------")
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("----------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Cap:", cap(mySlice), "Value:", mySlice[i])
	}
}
