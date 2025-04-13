package main

import (
	"fmt"
)

func main() {
	fmt.Println("functions in golang")
	greet()

	result := adder(2, 3)
	fmt.Println("answer is:", result)

	prores, mymsg := proAdder(1, 2, 3, 4)
	fmt.Println("answer is", prores)
	fmt.Println("pro message is ", mymsg)
}

func greet() {
	fmt.Println("hey there")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "hi pro result function"
}
