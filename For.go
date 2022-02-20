package main
// go run for.go

import "fmt"


func fors(){

	names := []string {"Carlos", "Fernando", "Mauricio", "Armando", "Jose"}

	for index, name := range names {
		// This doesnt Change
		name = "New"
		fmt.Printf("Student %v name: %v\n", (index + 1), name)	
	}

	fmt.Println("")

	for i := 0; i < len(names); i++ {
		// names[i] = "New"
		fmt.Printf("Student %v name: %v\n", (i + 1), names[i])
	}
	


}