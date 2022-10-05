package main

import "fmt"

func main() {
	a := struct {
		intero int
	}{
		intero: 5,
	}
	fmt.Println(a)
}
