package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Hi"] = "Oi"
	myGreeting["Bye"] = "Tchau"
	fmt.Println(myGreeting)
}
