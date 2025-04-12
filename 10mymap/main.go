package main

import (
	"fmt"
)

func main() {
	fmt.Print("map in golang")

	language := make(map[string]string)

	language["JS"] = "javascript"
	language["RB"] = "ruby"
	language["PY"] = "python"
	language["JA"] = "java"

	fmt.Println("list of all language:", language)
	fmt.Println("JS short for :", language["JS"])

	delete(language, "RB")
	fmt.Println("list of all language:", language)

	//loop are interesting in golang

	for _, value := range language {
		fmt.Printf("for key v, value is %v\n", value)
	}

}
