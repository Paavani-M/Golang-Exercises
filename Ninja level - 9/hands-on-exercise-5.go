package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var a int64

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&a, 1)
			atomic.LoadInt64(&a)
			fmt.Println(a)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value", a)

}

// Fix the race condition using package atomic
