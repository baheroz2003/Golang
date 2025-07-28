// This is the main package where the program starts.
package main

// Importing necessary packages:
import (
	"fmt"          // For printing output to the console
	"io/ioutil"    // For reading data (like HTTP response body)
	"net/http"     // For making HTTP requests (like GET, POST, etc.)
	"strings"      // For string manipulation (e.g., using a string builder)
)

// Entry point of the program
func main() {
	fmt.Println("Welcome to web verb video") // Print a welcome message
	PerformGetRequest()                      // Call the function to perform a GET request
}

// This function performs a GET request to a URL and processes the response
func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts/1." // The target URL

	// Make the GET request to the given URL
	response, err := http.Get(myurl)
	if err != nil {
		// If there is an error in making the request, stop the program
		panic(err)
	}
	defer response.Body.Close() // Ensure the response body is closed when function ends

	// Print the HTTP status code (e.g., 200 for OK)
	fmt.Println("Status code:", response.StatusCode)

	// Create a string builder to efficiently build the string response
	var responseString strings.Builder

	// Read the full content of the response body
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Write the content to the string builder and get the byte count
	byteCount, _ := responseString.Write(content)
	fmt.Println("byte count is:", byteCount) // Print how many bytes were written

	// Print the content length reported by the server (may be -1 if unknown)
	fmt.Println("the length is", response.ContentLength)

	// Print the response body as a string (direct)
	fmt.Println("Response body:", string(content))

	// Print the same response body from the string builder
	fmt.Println(responseString.String())
}
