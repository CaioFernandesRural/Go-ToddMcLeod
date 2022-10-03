package main

import "fmt"

func main() {
	m := map[string]int{
		"roger":   32,
		"claudin": 78,
		"L":       13,
	}
	fmt.Println(m)

	fmt.Println(m["claudin"])

	fmt.Println(m["gozo"])

	v, ok := m["gozo"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["claudin"]; ok {
		fmt.Println(v)
	}

	m["caio"] = 19

	for k, v := range m {
		fmt.Println(k, v)
	}
}
