package main

import "fmt"

func main() {

	// days := []string{"monday", "tuesdays", "wednesday", "thueday", "friday", "saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, days := range days {
	// 	fmt.Println("index is  and value is %v\n", days)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco

		}
		if rougueValue == 5 {
			rougueValue++
			break
		}
		fmt.Println(rougueValue)
		rougueValue++

	}

lco:
	fmt.Println("jumping on my head")
}
