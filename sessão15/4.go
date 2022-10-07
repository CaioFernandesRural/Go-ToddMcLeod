package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Nome: ", p.first, p.last)
	fmt.Println("Idade: ", p.age)
}

func main() {
	a := person{
		first: "caio",
		last:  "fernandes",
		age:   19,
	}
	a.speak()
}
