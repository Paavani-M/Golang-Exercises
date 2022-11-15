package main

import "fmt"

func main() {

	favSport := "hockey"

	switch favSport {
	case "cricket":
		fmt.Println("Inside cricket")
	case "football":
		fmt.Println("Inside football")
	case "hockey":
		fmt.Println("inside hockey")
	}
}

// Create a program that uses a switch statement with the switch expression specified as a
// variable of TYPE string with the IDENTIFIER “favSport”.
