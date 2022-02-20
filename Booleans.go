package main
// go run booleans.go

import "fmt"

func booleans(){
	
	names := []string {"Carlos", "Fernando", "Mauricio", "Armando", "Jose"}
	age:=45

	 if (age > 45){
		 fmt.Println("You are old")
	 } else if(age > 65){
		 fmt.Println("You are getting retired")
	 } else {
		 fmt.Println("Enjoy you youth")
	 }

	 for index, name := range names {
		 if (index == 1) {
			fmt.Printf("Not going to print %v\n", name)
			// TODO: CONTINUE RETURNS TO THE TOP LEVEL OF THE LOOP
			continue
		 }
		 if (index == 2) {
			// TODO: BREAKS THE LOOP
			break
		 }

		 fmt.Println(name)
		 
	 }
}