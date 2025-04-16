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
	// EncodeJson()
	// DecodeJson()
	DecodeJson()

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

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	 {
                "coursename": "angular bootcamp",
                "Price": 2999,
                "website": "lco.in",
                "welcome to tags[]": null,
        }
	
	`)

	var lcoCourses course

	checkJson := json.Valid(jsonDataFromWeb)

	if checkJson {
		fmt.Println("json was  valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("json was invalid")
	}

	var myOnineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnineData)
	fmt.Printf("%#v/n", myOnineData)

	for k, v := range myOnineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
