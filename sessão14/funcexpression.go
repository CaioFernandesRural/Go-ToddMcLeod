package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("func expression")
	}
	f()
	g := func(x int) {
		fmt.Println("o ano do grande irmão: ", x)
	}
	g(1984)
}
