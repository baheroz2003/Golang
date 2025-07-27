package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")
	// var ptr *int
	// fmt.Println("value of pointer is",ptr)
	//object that stores the memory address
	myNumber:=23
	var ptr=&myNumber
	//refrencing is used with &
	fmt.Println("Value of actual pointer is ",ptr)
	fmt.Println("Value of actual pointer is ",*ptr)
	*ptr=*ptr * 2
	fmt.Println("new value is:",myNumber)
}
