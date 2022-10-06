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
	fmt.Println("eu sou ", s.first, s.last)
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
	ag1.speak()
}
