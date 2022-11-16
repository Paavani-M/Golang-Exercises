package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	a := 1

	wg.Add(10)

	var m sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			m.Lock()
			b := a
			//	runtime.Gosched()
			b++
			a = b
			fmt.Println(a)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value", a)

}

// Fix the race condition you created in the previous exercise by using a mutex
