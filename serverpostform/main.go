package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	myurl := "http://localhost:8000/postform"

	formData := url.Values{}
	formData.Add("username", "baheroz")
	formData.Add("username", "zeya") // multiple values for same key
	formData.Add("email", "zeya@example.com")

	resp, err := http.PostForm(myurl, formData)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("✅ Form POST sent. Status:", resp.Status)
	fmt.Println(string(content))
}
