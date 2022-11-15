package main

import "fmt"

func main() {
	for i := 10; i < 100; i++ {
		fmt.Printf("When %v is dividied by %v gives remainder %v", i, 4, i%4)
		fmt.Println()
	}
}

// Print out the remainder (modulus) which is found for each number between 10 and 100 when it
// is divided by 4.
