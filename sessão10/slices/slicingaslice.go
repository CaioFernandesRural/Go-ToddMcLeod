package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[2:5])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

//slice agrupa valores do mesmo tipo
