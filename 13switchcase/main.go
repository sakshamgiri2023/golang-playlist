package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch  and case in golang")

	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and roll dice again")
	default:
		fmt.Println("fuck you bruhhh")
	}

}
