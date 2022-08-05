package main

import "fmt"


type student struct {
	rollNo int
	studentName string
	favSUBJECT string
}


func main()  {
	
	student1 := student{
		rollNo: 23,
		studentName: "manas",
		favSUBJECT: "computer science",
	}

	fmt.Println(student1)
}