package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

func (p *person) speak() {
	fmt.Println(p.first_name, p.last_name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first_name: "Paavani",
		last_name:  "M",
		age:        21,
	}

	saySomething(&p1)

	//saySomething(p1) shows error
}

// ● create a type person struct
// ● attach a method speak to type person using a pointer receiver
// ○ *person
// ● create a type human interface
// ○ to implicitly implement the interface, a human must have the speak method
// ● create func “saySomething”
// ○ have it take in a human as a parameter
// ○ have it call the speak method
// ● show the following in your code
// ○ you CAN pass a value of type *person into saySomething
// ○ you CANNOT pass a value of type person into saySomething
