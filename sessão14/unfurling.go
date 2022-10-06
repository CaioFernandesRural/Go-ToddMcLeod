package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(xi...)
	fmt.Println("total é ", sum)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
		fmt.Printf("adicionando %v ao total = %v\n", v, sum)
	}
	return sum
}

//fun (r receiver) identifier (parametros) (returns) { codigo }
