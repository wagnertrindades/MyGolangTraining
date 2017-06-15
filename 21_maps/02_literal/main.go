package main

import "fmt"

func main() {
	var myGreeting = map[string]string{}
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)

	myGreeting2 := map[string]string{
		"Hello": "Hi",
		"Bye":   "Good Bye",
	}

	fmt.Println(myGreeting2)
}
