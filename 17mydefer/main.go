/*
The defer statement in Go schedules a function call
to be executed after the surrounding function completes.
It's commonly used for cleanup actions, ensuring resources
are released or operations are finalized regardless of how the function exits,
whether through normal completion, an error, or a panic.

The deferred function's arguments are evaluated immediately
when the defer statement is encountered, but the function call itself
is executed after the surrounding function returns.
If multiple defer statements are present, they are
executed in LIFO (Last-In-First-Out) order,
meaning the last deferred call is executed first.
*/

package main

import "fmt"

func main() {
	myDefer()
	fmt.Println("welome to defer in golang")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello world")

}

func myDefer() {
	for i := 0; i < 5; i++ {
		fmt.Print(i)

	}
	fmt.Println()
}
