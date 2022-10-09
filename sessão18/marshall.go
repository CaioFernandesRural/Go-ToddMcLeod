package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Primeiro string
	Ultimo   string
	Idade    int
}

func main() {
	a := pessoa{
		Primeiro: "caio",
		Ultimo:   "fernandes",
		Idade:    19,
	}
	b := pessoa{
		Primeiro: "andrea",
		Ultimo:   "silveira",
		Idade:    48,
	}
	pessoas := []pessoa{a, b}
	fmt.Println(pessoas)

	bs, err := json.Marshal(pessoas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
