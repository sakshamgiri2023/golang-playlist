package main

import "fmt"

func main() {
	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "tomato"
	fruitlist[2] = "peach"

	fmt.Println("fruit list is: ", fruitlist)
	fmt.Println("fruit list is ", len(fruitlist))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("veg list is ", vegList)
}
