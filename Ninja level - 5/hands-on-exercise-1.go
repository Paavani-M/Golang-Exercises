package main

import "fmt"

type person struct {
	first_name          string
	last_name           string
	favorite_ice_creams []string
}

func main() {

	p1 := person{
		first_name:          "Paavani",
		last_name:           "M",
		favorite_ice_creams: []string{"vanilla", "chocolate", "strawberry"},
	}

	p2 := person{
		first_name:          "Bhars",
		last_name:           "S",
		favorite_ice_creams: []string{"hazelnut", "pista"},
	}

	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)
	for i, v := range p1.favorite_ice_creams {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first_name)
	fmt.Println(p2.last_name)
	for i, v := range p2.favorite_ice_creams {
		fmt.Println(i, v)
	}

}

// Create your own type “person” which will have an underlying type of “struct” so that it can store
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavors.
