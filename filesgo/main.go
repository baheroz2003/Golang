package main

import (
	"fmt"       // for printing output
	"io"        // for writing string to file
	"io/ioutil" // for reading file content
	"os"        // for file creation and handling
)

func main() {
	fmt.Println("files in golang") // print a heading

	content := "This needs to go in a file - LCO" // content to write into the file

	// Create the file (if file already exists, it will be overwritten)
	file, err := os.Create("./mygofile.txt")
	checkNilErr(err)         // check if file creation failed
	defer file.Close()       // make sure file is closed when function exits

	// Write the string content into the file
	length, err := io.WriteString(file, content)
	checkNilErr(err)         // check if write failed

	fmt.Println("length is", length) // print how many bytes were written

	// Now read the file content and print it
	readFile("./mygofile.txt")
}

// Function to read content of the file and print it
func readFile(filename string) {
	// Read entire file content into byte slice
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err) // check if read failed

	// Convert bytes to string and print
	fmt.Println("text data inside the file is:\n", string(databyte))
}

// Function to check and panic if error is not nil
func checkNilErr(err error) {
	if err != nil {
		panic(err) // stop execution and show error if any
	}
}
