package main

import "fmt"

func updateName(n *string){
		*n = "Updated String"
	}

func pointer(){
	name:= "Juan"

	n := &name

	fmt.Println(name)
	updateName(n)
	fmt.Println(name)


}