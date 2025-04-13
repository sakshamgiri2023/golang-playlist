package main

import (
	"fmt"
	"net/url"
)

const myurls string = "https:\\lco.dev:3000/learn?coursename=reactjs&paymentid=dbfsbfjSB"

func main() {
	fmt.Println("handling urls in golang")
	fmt.Println(myurls)

	//parsing
	result, _ := url.Parse(myurls)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	// fmt.Println(result.)
	// fmt.Println(result.Path)

	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("params is :", val)
	}

	partsOfUrls := &url.URL{
		Scheme:  "hhttps",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user=saksham",
	}

	anotherUrl := partsOfUrls.String()
	fmt.Println(anotherUrl)

}
