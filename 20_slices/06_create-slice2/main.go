package main

import "fmt"

func main() {
	mySlice := []string{}
	myOtherSlice := [][]string{}
	fmt.Println(mySlice)
	fmt.Println(myOtherSlice)
	fmt.Println(mySlice == nil)
	fmt.Printf("%T \n", mySlice)
	fmt.Println("Len:", len(mySlice), "Cap:", cap(mySlice))
}
