package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("First_name: %s\n", p.first)
	fmt.Printf("Last_name: %s\n", p.last)
	fmt.Printf("Age: %d", p.age)
}

func main() {
	p1 := person{
		first: "Paavani",
		last:  "M",
		age:   21,
	}

	p1.speak()

}

// ● Create a user defined struct with
// ○ the identifier “person”
// ○ the fields:
// ■ first
// ■ last
// ■ age
// ● attach a method to type person with
// ○ the identifier “speak”
// ○ the method should have the person say their name and age
// ● create a value of type person
// ● call the method from the value of type person
