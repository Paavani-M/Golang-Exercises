package main

import "fmt"

func main() {
	a := 10
	b := 12

	if a < b {
		fmt.Println("inside if")
	} else if a > b {
		fmt.Println("inside else if")
	} else {
		fmt.Println("inside else")
	}
}

// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
