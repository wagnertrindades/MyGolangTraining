package main

import "fmt"

func main() {

	myGreetigns := map[int]string{
		0: "Hi",
		1: "Hello",
		2: "Good bye!",
	}

	for key, value := range myGreetigns {
		fmt.Println(key, "-", value)
	}
}
