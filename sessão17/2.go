package main

import "fmt"

type person struct {
	nome string
}

func main() {
	a := person{
		nome: "caio",
	}
	fmt.Println(a)

	changeMe(&a)

	fmt.Println(a)
}

func changeMe(p *person) {
	(*p).nome = "roger"
}
