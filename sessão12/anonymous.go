package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
	}{
		first: "caio",
		last:  "fernandes",
	}

	fmt.Println(p1)

	fmt.Println(p1.first, p1.last)
}
