package main

import (
	"fmt"
	"sync"
)

func main() {
	con := 0

	gs := 50

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := con
			v++
			//runtime.Gosched()
			con = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("con: ", con)
}
