package main

import "fmt"

type person struct {
	first string
	last  string
}

type agente struct {
	person
	licensa bool
}

func (s agente) speak() {
	fmt.Println("eu sou agente: ", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("eu sou pessoa: ", p.first, p.last)
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("passei em bar: ", h.(person).first)
	case agente:
		fmt.Println("passei em bar: ", h.(agente).first)
	}
}

// fun (r receiver) identifier (parametros) (returns) { codigo }
func main() {
	ag1 := agente{
		person: person{
			first: "james",
			last:  "bond",
		},
		licensa: true,
	}
	p1 := person{
		first: "doutor",
		last:  "não",
	}
	ag1.speak()
	fmt.Println(p1)

	bar(ag1)
	bar(p1)
}
