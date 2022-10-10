package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Rotinas", runtime.NumGoroutine())

	var contador int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("contador: ", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Rotinas", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("contador: ", contador)
	fmt.Println("Rotinas", runtime.NumGoroutine())
}
