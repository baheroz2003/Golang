package main

import "fmt"

func main() {
	fmt.Println("welcome to array in go lang")
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "tomato"
	// fruitlist[2]="mango"
	fruitlist[3] = "orange"
	fmt.Println("fruit list is:", fruitlist)
	fmt.Println("fruit list is:", len(fruitlist))
	var veglist=[5]string{"potato","beans","mushroom"}
	fmt.Println("vegy list is:",len(veglist))
}
