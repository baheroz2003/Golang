package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to video on slices")

	// Creating a slice
	var fruitList = []string{"apple", "mango"}

	// Printing the type
	fmt.Printf("type of fruitlist is %T\n", fruitList)

	// Appending to the slice
	fruitList = append(fruitList, "banana")
	fmt.Println("fruitList:", fruitList) // Output: [apple mango banana]

	// Removing the first element (slice from index 1 to end)
	// fruitList = append(fruitList[1:])
	fmt.Println(fruitList) // Output: [mango banana]
	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 234
	highScore[2] = 234
	highScore[3] = 234
	//do not add like u did above using like a[4]
	// Append 255 as 5th element
	highScore = append(highScore, 255, 666)
	//it will realloacte the memory
	fmt.Println(highScore) // Output: [234 234 234 234 255]
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	//how to remove the element in slice
	var courses=[]string{"reactjs","js","swift","go","cpp"}
	fmt.Println(courses)
	var index int=2
	courses=append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
	//... used to allow a function to accept a variable no of arguments:variadic function
}
