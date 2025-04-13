// function is uesd for general purpose
// but if talk abt class(structs in golang) in method is taken in consideration which is nothing
// but  a reference to funtion itself

// if variable first letter is capital
// then it is public in nature otherwise not

package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh detail are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Printf("Name is _ and email is %v.\n", hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("email of this user is: ", u.Email)
}
