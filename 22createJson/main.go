package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"welcome to tags[]"`
}

func main() {
	fmt.Println("creating a json in golang")
	EncodeJson()

}

func EncodeJson() {

	lcoCourses := []course{
		{"reactjs bootcamp", 299, "lco.in", "abc123", []string{"fd"}},
		{"mern bootcamp", 2199, "lco.in", "abc123", []string{"fs"}},
		{"angular bootcamp", 2999, "lco.in", "abc123", nil},
	}

	finalJson, _ := json.MarshalIndent(lcoCourses, "", "\t")
	fmt.Printf("%s\n", finalJson)
}
