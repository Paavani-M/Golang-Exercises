package main

import (
	"fmt"
)

func main() {

	m := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				m <- j
			}
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i, <-m)
	}

	fmt.Println("EXIT")

}

// ● write a program that
// ○ launches 10 goroutines
// ■ each goroutine adds 10 numbers to a channel
// ○ pull the numbers off the channel and print them
