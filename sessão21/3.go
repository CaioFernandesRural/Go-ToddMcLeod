package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	con := 0

	gs := 50

	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := con
			v++
			runtime.Gosched()
			con = v
		}()
		wg.Done()
	}
	wg.Wait()
	fmt.Println("con: ", con)
}
