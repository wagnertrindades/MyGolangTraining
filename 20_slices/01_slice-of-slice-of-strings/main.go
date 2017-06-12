package main

import "fmt"

func main() {

	records := make([][]string, 0)
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Wagner"
	student1[1] = "Trindade"
	student1[2] = "100.00"
	student1[3] = "80.00"
	// store the record
	records = append(records, student1)

	// student 2
	student2 := make([]string, 4)
	student2[0] = "Steve"
	student2[1] = "Vai"
	student2[2] = "90.00"
	student2[3] = "70.00"
	// store the record
	records = append(records, student2)

	fmt.Println(records)
}
