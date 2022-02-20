package main

import "fmt"

	func updateName(n string){
		n = "Updated String"
	}

func passByValue(){
	// TODO: The value is not changed dirrectly in the function
	// 1 group strings, ints, bools, floats, arrays, structs

	// TODO: The value does change
	// 2 group maps, slices, functions
	 
	name:="Mauricio"

	updateName(name)

	fmt.Println(name)
	// -> Mauricio 

	// TODO: To update function you should return the type ans store it in the value
	// func updateName(n string) string{
	// 	n = "Updated String"
	// }
}