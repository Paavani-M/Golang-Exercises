package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))

}

// Instead of using the blank identifier, make sure the code is checking and
// handling the error.
