package main

import "fmt"

func main() {
	logincnt := 23
	var result string

	if logincnt > 10 {
		result = "regular user"
	} else if logincnt < 10 {
		result = "watch out"
	} else {
		result = "equal"
	}

	fmt.Println(result)
	if 9%2==0{
		fmt.Println("number is even")
	}else{
		fmt.Println("number is odd")
	}
	if num:=3; num<10{
		fmt.Println("num is less than 10")
	}else{
		fmt.Println("num is not less than 10")
	}
}
