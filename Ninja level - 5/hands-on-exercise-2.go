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

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for _, v := range m {
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for i1, v1 := range v.favorite_ice_creams {
			fmt.Println(i1, v1)
		}
	}
}

// Take the code from the previous exercise, then store the values of type person in a map with the
// key of last name. Access each value in the map. Print out the values, ranging over the slice.
