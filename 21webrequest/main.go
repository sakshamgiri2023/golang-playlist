package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("get web request in golang")
	PerformGetRequest()
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
