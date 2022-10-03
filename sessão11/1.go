package main

import "fmt"

func main() {
	var arr [5]int

	for i := range arr {
		arr[i] = i + 1
		fmt.Println(arr[i])
	}
	fmt.Printf("%T\n", arr)

}
