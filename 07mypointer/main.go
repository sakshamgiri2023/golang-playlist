package main

import "fmt"

func main() {
	fmt.Println("notes of the pointer on golang")

	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("value of actual is", ptr)
	fmt.Println("value of actual is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("new value is ", myNumber)

}
