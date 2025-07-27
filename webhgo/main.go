package main // This defines the main package — required for standalone Go applications

import (
	"fmt"       // Provides formatted I/O operations like Println, Printf
	"io/ioutil" // For reading data from the response body
	"net/http"  // For making HTTP requests
)

func main() {
	// Send an HTTP GET request to the specified URL
	response, err := http.Get("https://www.codecademy.com")
	if err != nil {
		// If there's an error while making the request, panic and stop execution
		panic(err)
	}

	// Ensure the response body is closed after function completes — prevents resource leaks
	defer response.Body.Close()

	// Print the type of the response object (it's *http.Response)
	fmt.Printf("response is of type: %T\n", response)

	// Read the entire response body (i.e., HTML content) into a byte slice
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// Panic if there's an error while reading the body
		panic(err)
	}

	// Convert byte slice to a string to view the HTML content as readable text
	content := string(databytes)

	// Print the content (HTML source code of the webpage)
	fmt.Println(content)
}
