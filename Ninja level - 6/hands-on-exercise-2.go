package main

import "fmt"

func main() {

	defer foo()

	bar()

}

func foo() {
	fmt.Println("Inside foo function")
}

func bar() {
	fmt.Println("Inside bar function")
}

// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
