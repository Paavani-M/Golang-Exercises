package main

import "fmt"

type person struct {
	first_name string
}

func changeMe(p *person) {
	(*p).first_name = "Bharathi"
}

func main() {
	p1 := person{
		first_name: "Paavani",
	}
	fmt.Println(p1.first_name)

	changeMe(&p1)

	fmt.Println(p1.first_name)

}

// ● create a person struct
// ● create a func called “changeMe” with a *person as a parameter
// ○ change the value stored at the *person address
// ● create a value of type person
// ○ print out the value
// ● in func main
// ○ call “changeMe”
// ● in func main
// ○ print out the value
