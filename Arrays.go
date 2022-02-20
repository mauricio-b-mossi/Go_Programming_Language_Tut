package main
// go run arrays.go

import "fmt"

// TODO: SLICE METHODS ARE FUNCTIONS
func arrays(){
	// TODO: ARRAYS
	// var ages [3]int = [3]int{1, 2, 3}
	var ages = [3]int{1, 2, 3}
	names := [4]string{"Valeria", "Mauricio", "Ana", "Cristina"}

	// TODO: SLICES: ArrayList in java
	scores := []int{}
	scores = append(scores, 5, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10)
	scores[0] = 20

	// slice ranges (same as py)
	scoreRange := scores[5:6]
	scoreRangeTwo := scores[:3]

	

	fmt.Println(len(ages))
	fmt.Printf("names: %q \n", names )
	fmt.Printf("scores: %v \n", scores )
	fmt.Printf("scoresMiddle: %v \n", scoreRange )
	fmt.Printf("scoresTwo: %v \n", scoreRangeTwo )
	
}