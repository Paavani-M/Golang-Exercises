package main

import "fmt"

func main() {

	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	j := [][]string{a, b}

	for _, v := range j {
		fmt.Println(v)
		for _, v1 := range v {
			fmt.Println(v1)
		}

	}

}

// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
// slice:
// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."
// Range over the records, then range over the data in each record.
