package main

import "fmt"

var a int

type hotdog int

var b hotdog = 43

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
