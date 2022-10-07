package main

import "fmt"

func main() {
	a, b := bar()
	fmt.Println(foo(), a, b)
}

func foo() int {
	return 19
}
func bar() (int, string) {
	return 19, "dezenove"
}
