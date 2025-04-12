package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15-04-2006 Monday"))

	createdDated := time.Date(2020, time.June, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDated)
	fmt.Println(createdDated.Format("01-02-2006 Monday"))

}
