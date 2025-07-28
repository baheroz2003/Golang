package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"coursename": "Let's go with Golang",
		"price": 0,
		"platform": "lco.in"
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code:", response.StatusCode)
	// fmt.Println("Response Body:", string(body))
	fmt.Println(string(content))
}
