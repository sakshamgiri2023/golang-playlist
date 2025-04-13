package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("lco web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response is of type: %T\n", response)
	defer response.Body.Close() // caller's responsibility to clsoe the connectivity

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}
