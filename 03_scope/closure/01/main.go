package main

import "fmt"

func main() {
	x := 42
	{
		fmt.Println(x)
		y := "This is inner scope"
		fmt.Println(y)
	}
	// fmt.Println(y) // outside of scope y
}
