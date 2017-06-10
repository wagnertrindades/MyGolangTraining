package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Wagner"]) // 44
}

func changeMe(z map[string]int) {
	z["Wagner"] = 44
}
