package main

import "fmt"

func main() {
	mySlice := make([]string, 38, 100)
	myOtherSlice := make([][]string, 38)
	fmt.Println(mySlice)
	fmt.Println(myOtherSlice)
	fmt.Println(mySlice == nil)
	fmt.Println("Len:", len(mySlice), "Cap:", cap(mySlice))
	fmt.Println("Len:", len(myOtherSlice), "Cap:", cap(myOtherSlice))
}
