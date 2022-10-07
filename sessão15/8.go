package main

import "fmt"

func main() {
	f := a()

	f()
}
func a() func() {
	return func() {
		fmt.Println("aaaa")
	}
}
