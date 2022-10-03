package main

import "fmt"

func main() {
	m := map[string]int{
		"roger":   32,
		"claudin": 78,
		"L":       13,
	}
	fmt.Println(m)

	delete(m, "L")

	fmt.Println(m)
}
