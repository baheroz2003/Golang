package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()
	greeterTwo()
	result:=adder(2,5)
	fmt.Println(result)
	prores,msg:=proadder(2,5,8,7)
	fmt.Println(prores)
	fmt.Println(msg)
}
//... variadic function which can accept variable no of arguments
func adder(val1 int,val2 int) int{
	return val1+val2
}
//that is call function signature taht is what type of value u want to return
func proadder(values ...int)(int,string){
	total:=0
	for _,val:=range values{
		total+=val
	}
	return total,"hi you get result"
}
func greeter() {
	fmt.Println("namastey from golang")
}
func greeterTwo() {
	fmt.Println("another method")
}
//you can define func without name that is annonymous fucntion exists

