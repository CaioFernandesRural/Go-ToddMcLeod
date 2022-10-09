package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Primeiro string `json:"Primeiro"`
	Ultimo   string `json:"Ultimo"`
	Idade    int    `json:"Idade"`
}

func main() {
	s := `[{"Primeiro":"caio","Ultimo":"fernandes","Idade":19},{"Primeiro":"andrea","Ultimo":"silveira","Idade":48}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var pessoas []pessoa

	err := json.Unmarshal(bs, &pessoas)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pessoas)
}
