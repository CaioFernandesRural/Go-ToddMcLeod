package main

import "fmt"

type pessoa struct {
	first string
	last  string
	favs  []string
}

func main() {
	a := pessoa{
		first: "caio",
		last:  "fernandes",
		favs:  []string{"passas ao rum", "flocos"},
	}
	b := pessoa{
		first: "andrea",
		last:  "fernandes",
		favs:  []string{"baunilha", "napolitano"},
	}
	fmt.Println(a.first, a.last)
	for _, v := range a.favs {
		fmt.Println(v)
	}
	fmt.Println(b.first, b.last)
	for _, v := range b.favs {
		fmt.Println(v)
	}
}
