package main

//go run standardlibary.go

import (
	"fmt"
	"sort"
	"strings"
)



func standardlibary(){

	// TODO: STRINGS

	greetings := "Hello there"
	changedToA := strings.ReplaceAll(greetings, "e", "a")
	contains := strings.Contains(greetings, "there")
	upper := strings.ToUpper(greetings)
	// split := strings.Split(greetings, "")


	fmt.Println(contains)
	fmt.Println(changedToA)
	fmt.Println(upper)
	// for i := 0; i < len(split); i++ {
	// 	fmt.Println(split[i])
	// }

	// TODO: SORT the
	ages:= []int{20, 34, 53, 43, 76, 23, 546, 123 ,678, 23, 345 ,123 ,56 ,24}
	sort.Ints(ages)

	fmt.Println(ages)

	


}