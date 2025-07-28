package main

import (
	"encoding/json" // Package to handle JSON encoding/decoding
	"fmt"           // Package for formatted I/O
)

// Define a `course` struct with JSON tags for custom key names and formatting
type course struct {
	Name     string   `json:"coursename"`     // Change the key name to "coursename" in JSON
	Price    int      `json:"price"`          // Normal int field mapped to "price"
	Platform string   `json:"website"`        // Corrected spelling from "Platfrom" to "Platform"; mapped to "website"
	Password string   `json:"-"`              // "-" tells Go to skip this field in JSON output
	Tags     []string `json:"tags,omitempty"` // "omitempty" skips this field if it's nil or empty
}

func main() {
	fmt.Println("Welcome to JSON video") // Simple output
	// EncodeJson()                         // Call the function to encode struct into JSON
	DecodeJson()
}

func EncodeJson() {
	// Sample data: slice of course structs
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"NodeJS Bootcamp", 199, "LearnCodeOnline.in", "abcd123", []string{"web-dev", "js"}},
		{"AJS Bootcamp", 299, "LearnCodeOnline.in", "abcd1e23", nil}, // Tags is nil → will be omitted due to "omitempty"
	}

	// MarshalIndent: converts Go object to pretty-printed JSON
	// "" → no prefix for each line, "\t" → use tab character for indentation
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err) // Stop execution if an error occurs during JSON conversion
	}

	// Print the final JSON output as a string
	fmt.Printf("%s\n", finalJson)
}
func DecodeJson() {
	// Raw JSON data (simulating a response from an API or file)
	jsonData := []byte(`
			{
				"coursename": "ReactJS Bootcamp",
				"Price": 299,
				"website": "LearnCodeOnline.in",
				"tags": ["web-dev", "js"]
			}
	`)
	var lcoCourse course
	checkValid :=json.Valid(jsonData)
	if checkValid{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData,&lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	}else{
		fmt.Println("json not valid")
	}
	//some cases where you want to add in key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData,&myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)
	for k,v :=range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is :%T\n",k,v,v)
	}
}