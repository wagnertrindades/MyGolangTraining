package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Hi",
		1: "Hello",
		2: "Bye",
		3: "Good bye!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 2)
	fmt.Println(myGreeting)
}
