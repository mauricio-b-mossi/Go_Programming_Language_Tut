package main

import "fmt"

func maps(){

	// Key value pairs int the brackets [type of key] after brackets type of value
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee" : 3.55,
	}

	phoneNumbers := map[int]string{
		25190087 : "casa",
		78607921 : "mama",
		78609285 : "papa",
	}

	fmt.Println(phoneNumbers)

	for k, v := range menu {
		fmt.Printf("key: %v value: %v \n",k, v)
	}
}