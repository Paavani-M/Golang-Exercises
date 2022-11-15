package main

import "fmt"

func main() {
	a := sum([]int{1, 2, 3, 4, 5, 6})

	b := add(a, 9)

	fmt.Printf("Inside add function %d\n", b)
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}

	fmt.Printf("Inside sum function: %d\n", sum)

	return sum

}

func add(a int, b int) int {
	return a + b
}

// A “callback” is when we pass a func into a func as an argument. For this exercise,
// ● pass a func into a func as an argument
