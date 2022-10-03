package main

import "fmt"

func main() {
	x := make([]int, 10, 12) //tipo, tamanho alocado, tamanho pré-alocado máximo
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3, 5, 6) //se estoura, dobra o máximo

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

//slice agrupa valores do mesmo tipo
