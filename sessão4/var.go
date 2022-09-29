package main

import "fmt"

var y = 100 + 24 //var para variáveis globais

func main() {
	x := 42 //declara e dá o tipo automaticamente
	fmt.Println(x)

	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println("dnv, ", y)
}
