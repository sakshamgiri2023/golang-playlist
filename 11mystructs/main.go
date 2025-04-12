package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh detail are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
