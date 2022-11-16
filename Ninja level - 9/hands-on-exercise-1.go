package main

import (
	"fmt"
	"runtime"
	"sync"
)

var a sync.WaitGroup

func main() {
	fmt.Println("Initial CPU", runtime.NumCPU())
	fmt.Println("Initial GR", runtime.NumGoroutine())

	fmt.Println("A")

	a.Add(2)
	go foo()

	go bar()

	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid GR", runtime.NumGoroutine())

	a.Wait()

	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End GR", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("B")
	a.Done()
}

func bar() {
	fmt.Println("C")
	a.Done()
}

// ● in addition to the main goroutine, launch two additional goroutines
// ○ each additional goroutine should print something out
// ● use waitgroups to make sure each goroutine finishes before your program exist
