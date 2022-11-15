package main

import "fmt"

func main() {

	func() {
		fmt.Println("inside anonymous function")
	}()

}

// ● Build and use an anonymous func
