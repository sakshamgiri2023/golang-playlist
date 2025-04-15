package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("get web request in golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("content length is:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
	{
	"coursename": "let's go with golang",
	"price": 0,
	"platform":"lcu"
	} `)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	data := url.Values{}

	data.Add("firstname", "saksham")
	data.Add("lastname", "Giri")
	data.Add("email", "giri@gmail.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
