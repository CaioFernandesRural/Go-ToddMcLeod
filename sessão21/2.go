package main

import "fmt"

type person struct {
	nome string
}

func (p *person) speak() {
	fmt.Println("opa, ", p.nome)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	a := person{"roger"}

	saySomething(&a)
	//saySomething(a) interface só reconhece com receiver ponteiro
}
