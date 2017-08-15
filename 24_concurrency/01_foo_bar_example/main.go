package main

import "fmt"

/*
	Not work because in this code have 3 go routines: main, foo, bar
	And main is executed more fast in concurrency
*/
func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i <= 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i <= 45; i++ {
		fmt.Println("Bar:", i)
	}
}
