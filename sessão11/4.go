package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arr = append(arr, 11)
	fmt.Println(arr)

	arr = append(arr, 12, 13, 14)
	fmt.Println(arr)

	y := []int{15, 16, 17}
	arr = append(arr, y...)
	fmt.Println(arr)
}
