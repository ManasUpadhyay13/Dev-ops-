package main

import "fmt"

func main(){


	// consider this array

	var array [10]int = [10]int{1 ,2 ,3 ,4 ,5 ,6 ,7 ,8 ,9 ,10}

	fmt.Println(array)

	// now slice code below 

	slice1 := array[:]  // this simply copies all the  elements of  the array into slice1

	slice1[6] = 45 // now since here i  have changed the value of the 6th element to 45
			      // it value gets changed all over the program as the slice is referenced to array's address

	slice2 := array[2:]   // elements from 3 to end

	slice3 := array[:8] // elements from from stating to 8th element

	slice4 := array[4:9]  // elements from from 5th to 9th

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

}