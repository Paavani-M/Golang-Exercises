package main

import "fmt"

func main() {

	a := struct {
		first_name string
		last_name  string
		age        int
	}{
		first_name: "Paavani",
		last_name:  "M",
		age:        21,
	}

	fmt.Printf("first_name: %s\n", a.first_name)
	fmt.Printf("last_name: %s\n", a.last_name)
	fmt.Printf("age: %d\n", a.age)

}

// Create and use an anonymous struct.
