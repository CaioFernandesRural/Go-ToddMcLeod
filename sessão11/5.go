package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	arr = append(arr[:3], arr[6:]...)
	fmt.Println(arr)
}
