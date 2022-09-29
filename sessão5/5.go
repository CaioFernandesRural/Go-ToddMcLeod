package main

import "fmt"

type banana int

var x banana

var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 19
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
