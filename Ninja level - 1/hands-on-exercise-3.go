package main

import "fmt"

var x = 32
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}

// 1. At the package level scope, assign the following values to the three variables
// a. for x assign 42
// b. for y assign “James Bond”
// c. for z assign true
// 2. in func main
// Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 26
// a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
// returned value of TYPE string using the short declaration operator to a
// VARIABLE with the IDENTIFIER “s”
// b. print out the value stored by variable “s”
