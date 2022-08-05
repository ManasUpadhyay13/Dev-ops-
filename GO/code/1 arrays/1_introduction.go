// this is an introduction of arrays in GO language.

package main

import "fmt"


func main()  {

	// in this method we declar an array and then immediately initialize the array with values

	tempArray := [...]int{1 ,2 ,3 ,4 ,5 ,6 ,7 ,8 ,9}
	fmt.Printf("%v\n", tempArray)


	// how to just declare the array , withour intitialization

	var tempArray2 [4]int 
	fmt.Printf("%v\n", tempArray2)


	// to find the length of the array

	fmt.Printf("%v\n", len(tempArray2))
} 