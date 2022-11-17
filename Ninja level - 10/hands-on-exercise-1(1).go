package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

// ● get this code working:
// ○ using buffered channel
