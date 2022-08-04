package main

import "fmt"

func main() {
	// in this we are going to learn about make method

	// make generally works two ways , one with two parameters and second one with three parameters

	// two parameters

	slice1 := make([]int , 10)
	fmt.Println( slice1 , len(slice1)  , cap(slice1))

	// this above was the fiest method in which the fiest parameter
	// that we pass in is the datatype of the array and the second one
	// is the length of the slice


	// two parameters

	slice2 := make([]int , 10 , 100)
	fmt.Println(slice2 , len(slice2) , cap(slice2))

	// this is the second method in which we pass 3 parameters 
	// the first one is the type of the array 
	// the second one is the length of the slice
	// and the third one is the capacity of the slice 



	//  ABOUT APPEND FUNCITOIN

	// append function is used to interst values in the slice
	// just like the make funciton the append funciton also take 
	// 2 and 3 parameters

	

	// 2 parameters append function

	slice1 = append(slice1, 2)
	fmt.Println(slice1)

	// 3 or more parameters append function

	slice2 = append(slice2, 3 , 4 , 5 , 6 , 7)
	fmt.Println(slice2)

	// in this operation we are appending more than 2 parameters at
	// a sigle time

}
