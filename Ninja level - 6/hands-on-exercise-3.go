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
