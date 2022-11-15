package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("hi")
	}

	a()
	b := a

	b()

}

// ● Assign a func to a variable, then call that func
