package main // defines the main package – required for executable Go programs

import (
	"fmt"    // for formatted I/O operations (like Println)
	"net/url" // provides functions for parsing and building URLs
)

// constant URL string to be parsed and analyzed
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymenat"

func main() {
	// Print welcome message
	fmt.Println("welcome to handling urls in golang")

	// Print the original URL
	fmt.Println(myurl)

	// Parse the URL string into a structured URL object
	result, _ := url.Parse(myurl)

	// Print different parts of the parsed URL
	fmt.Println("Scheme:", result.Scheme)     // https
	fmt.Println("Host:", result.Host)         // lco.dev:3000
	fmt.Println("Path:", result.Path)         // /learn
	fmt.Println("Port:", result.Port())       // 3000
	fmt.Println("RawQuery:", result.RawQuery) // coursename=reactjs&paymenat

	// Extract the query parameters into a map[string][]string
	qparams := result.Query()
	fmt.Printf("The type of query params is: %T\n", qparams) // url.Values (i.e., map[string][]string)

	// Access a specific query parameter using the key "coursename"
	fmt.Println(qparams["coursename"]) // [reactjs]

	// Loop through all query parameters and print key-value pairs
	for key, val := range qparams {
		fmt.Println("Param:", key, "Value:", val)
	}
	// Constructing a new URL manually using url.URL struct
	partsOfUrl := &url.URL{
		Scheme:   "https",           // protocol
		Host:     "lco.dev",         // domain
		Path:     "/tutcss",         // path
		RawQuery: "user=hitesh",     // query string (after ?)
	}
	// Convert the URL object to a string
	anotherURL := partsOfUrl.String()
	// Print the newly constructed URL
	fmt.Println("Constructed URL:", anotherURL) // https://lco.dev/tutcss?user=hitesh
}
