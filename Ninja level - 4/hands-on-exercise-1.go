package main

import "fmt"

func main() {
	a := [5]int{11, 12, 13, 14, 15}

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)
}

// ● Using a COMPOSITE LITERAL:
// ● create an ARRAY which holds 5 VALUES of TYPE int
// ● assign VALUES to each index position.
// ● Range over the array and print the values out.
// ● Using format printing
// ○ print out the TYPE of the array
