package main

import "fmt"

func main() {
	n := fatorial(6)
	fmt.Println(n)

	n1 := loopfac(6)
	fmt.Println(n1)
}
func fatorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}
func loopfac(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
