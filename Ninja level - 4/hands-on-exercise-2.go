package main

import "fmt"

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 89}

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", a)

}

// ● Using a COMPOSITE LITERAL:
// ● create a SLICE of TYPE int
// ● assign 10 VALUES
// ● Range over the slice and print the values out.
// ● Using format printing
// ○ print out the TYPE of the slice
