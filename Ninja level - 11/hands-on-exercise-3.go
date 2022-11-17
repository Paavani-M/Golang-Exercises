package main

import "fmt"

type customErr struct {
	first string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("There is an error %v", ce.first)
}

func foo(e error) {
	fmt.Println("Inside foo", e, "\n")
}

func main() {
	c := customErr{
		first: "Paavani",
	}

	foo(c)
}

// Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
// has a value of type error as a parameter. Create a value of type “customErr” and pass it into
// “foo”.
