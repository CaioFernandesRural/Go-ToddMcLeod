package main

import "fmt"

func main() {
	a := []int{16, 61, 85, 56, 26, 2, 86, 4}

	fmt.Println(foo(a...))
	fmt.Println(bar(a))
}
func foo(x ...int) int {
	som := 0

	for _, v := range x {
		som += v
	}
	return som
}
func bar(x []int) int {
	som := 0

	for _, v := range x {
		som += v
	}
	return som
}
