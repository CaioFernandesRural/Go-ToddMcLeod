package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)

}
