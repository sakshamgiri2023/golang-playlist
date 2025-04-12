package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to notes of slices")

	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)

	fruitList = fruitList[1:]
	fmt.Println(fruitList)

	highSore := make([]int, 5)

	highSore[0] = 1244
	highSore[1] = 1234
	highSore[2] = 1245
	highSore[3] = 1245
	highSore[4] = 1245

	highSore = append(highSore, 555, 666, 777, 321)

	fmt.Println(highSore)

	// fmt.Println(sort.IntsAreSorted(highSore))
	sort.Ints(highSore)
	fmt.Println(highSore)

}
