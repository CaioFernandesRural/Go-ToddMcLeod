package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Rotinas", runtime.NumGoroutine())

	contador := 100

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := contador
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Rotinas", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("contador: ", contador)
	fmt.Println("Rotinas", runtime.NumGoroutine())
}
