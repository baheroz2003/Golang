// main2.go

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	//see the err not equal to nil wala
}
