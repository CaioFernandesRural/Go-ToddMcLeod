package main

import "fmt"

func main() {

	defer b()

	fmt.Println("a")
}
func b() {
	fmt.Println("b")
}
