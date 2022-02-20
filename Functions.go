package main

// go run functions.go

import (
	"fmt"
	"math"
	"strings"
)

// TODO: nested functions write (var), func((type))\
// Return type used in the same way as TS
func individualGreeting(name string){
	fmt.Printf("Hi there %v\n", name)
}

func groupGreeting(names []string, f func(string)){
	for _, name := range names {
		f(name)
	
	}
}

func circleArea(r float64)float64 {
	return math.Pi * r * r
}

func getInitials(name string) (string, string){
	s := strings.ToUpper(name)
	separated := strings.Split(s, " ")

	var initials []string
	for _, v := range separated{
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	
	return initials[0], "_"
	
}

func returnInitials(name string) (string, string){
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func functions(){

	names := [4]string{"Valeria", "Mauricio", "Ana", "Cristina"}
	
	individualGreeting("Mauricio")

	groupGreeting(names[:],individualGreeting)

	area := circleArea(10)
	fmt.Printf("%0.2f\n", area)

	fistInitial, secondInitial := returnInitials("Fernando Hernandez")

	fmt.Println(fistInitial, secondInitial)

}