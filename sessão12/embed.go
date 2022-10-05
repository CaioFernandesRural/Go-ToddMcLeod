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

func main() {
	ag1 := agente{
		person: person{
			first: "caio",
			last:  "fernandes",
		},
		licensa: true,
	}
	p2 := person{
		first: "andrea",
		last:  "fernandes",
	}

	fmt.Println(ag1)
	fmt.Println(p2)

	fmt.Println(ag1.first, ag1.last, ag1.licensa)
	fmt.Println(p2.first, p2.last)
}
