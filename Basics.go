package main
// go run basics.go

// to print msgs && format strings
import "fmt"

// Main is same as java
func basics(){
	// Strings
	var nameOne string = "Mauricio"
	nameTwo:= "Benjamin"

	// ints: can be alocated specific bit size (int(8, 16, 32, 64))
	var numOne int8 = 1
	numTwo:=2
	// means only positive
	var numThree uint = 25

	// float: contains decimal point
	var numFloatOne float32 = 25.98
	var numFloatTwo float64 = 1231.12

	// Sprintf
 var fullName =fmt.Sprint(nameOne, nameTwo) 

// String formatting is the same as with python %v: val, %q: "_", %T: typeof, %(0.2)f: floats
	fmt.Println("My name is:",nameOne, nameTwo)
	fmt.Printf("numOne %v numTwo %v numThree %v \n", numOne, numTwo, numThree)
	fmt.Println(numFloatOne, numFloatTwo)
	fmt.Printf("The saved string is: %v", fullName)


}