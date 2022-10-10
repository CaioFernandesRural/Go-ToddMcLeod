package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	fmt.Println("a")

	go func() {
		fmt.Println("b")
		wg.Done()
	}()
	go func() {
		fmt.Println("c")
		wg.Done()
	}()
	wg.Wait()
}
