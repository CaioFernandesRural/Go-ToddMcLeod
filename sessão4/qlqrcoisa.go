package main

import "fmt"

func main() {
	fmt.Println("Hello homie")

	foo()

	fmt.Print("otracoisa")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("foo faiters")
}
func bar() {
	fmt.Println("cabo")
}

//demonstação de fluxo de controle
