package main

import (
	"fmt"
)

func  main() {
	// first way to create a map

	// map keywork then the type of the key and then the type of the value

	tempMap := map[string]int{
		"manas" : 100 ,
		"lucky" : 89 ,
		"vishal" : -12,
	}

	fmt.Println(tempMap)
	fmt.Println(tempMap["manas"])

	// adding values to the map

	tempMap["scout"] = 234
	fmt.Println(tempMap)
}