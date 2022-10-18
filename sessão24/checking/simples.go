package main

import "fmt"

func main() {
	n, err := fmt.Println("Coe")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
