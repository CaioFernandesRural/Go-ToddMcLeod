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
		last:  "silveira",
		favs:  []string{"baunilha", "napolitano"},
	}
	m := map[string]pessoa{
		a.last: a,
		b.last: b,
	}
	for _, v := range m {
		fmt.Println(v.first, v.last)
		for _, v := range v.favs {
			fmt.Println(v)
		}
	}
}
