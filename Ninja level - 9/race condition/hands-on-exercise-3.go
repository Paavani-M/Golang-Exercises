package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	a := 1

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			b := a
			runtime.Gosched()
			b++
			a = b
			fmt.Println(a)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value", a)

}

// ● Using goroutines, create an incrementer program
// ○ have a variable to hold the incrementer value
// ○ launch a bunch of goroutines
// ■ each goroutine should
// ● read the incrementer value
// ○ store it in a new variable
// ● yield the processor with runtime.Gosched()
// ● increment the new variable
// ● write the value in the new variable back to the incrementer
// variable
// ● use waitgroups to wait for all of your goroutines to finish
// ● the above will create a race condition.
// ● Prove that it is a race condition by using the -race flag
