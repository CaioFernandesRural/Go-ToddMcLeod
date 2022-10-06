package main

import "fmt"

func main() {
	defer foo()
	bar()
} //defer executada no fim da função apenas
func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
