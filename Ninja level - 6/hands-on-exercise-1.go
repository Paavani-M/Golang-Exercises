package main

import "fmt"

func main() {

	x := []int{4, 3, 2, 7, 8}
	y := foo(x...)

	fmt.Printf("Sum(Variadic parameter): %d\n", y)
	z := bar(x)

	fmt.Printf("Sum([]int): %d\n", z)

}

func foo(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}

	return sum

}

func bar(b []int) int {
	sum := 0
	for _, v := range b {
		sum += v
	}

	return sum
}

// ● create a func with the identifier foo that
// ○ takes in a variadic parameter of type int
// ○ pass in a value of type []int into your func (unfurl the []int)
// ○ returns the sum of all values of type int passed in
// ● create a func with the identifier bar that
// ○ takes in a parameter of type []int
// ○ returns the sum of all values of type int passed in
