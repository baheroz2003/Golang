package main

import (
	"bufio"  // Buffered I/O package, line-by-line input lene ke liye
	"fmt"    // Formatting package, print karne ke liye
	"os"     // OS package, input/output stream ke liye
)

func main() {
	// Welcome message print karega
	fmt.Println("Welcome to input")

	// Reader create kar rahe hain jo standard input (keyboard) se line read karega
	reader := bufio.NewReader(os.Stdin)

	// Prompt message user ke liye
	fmt.Println("Enter the rating for our pizza:")

	// User se ek line input le rahe hain (jab tak Enter press na ho)
	input, err := reader.ReadString('\n')

	// Agar koi error aayi input lene mein toh print karo
	if err != nil {
		fmt.Println("Error reading input:", err)
	} else {
		// Agar input sahi mila toh user ko thank you bolenge
		fmt.Println("Thanks for rating,", input)
	}
}
