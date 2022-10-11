package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var con int64

	gs := 50

	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&con, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("con: ", con)
}
