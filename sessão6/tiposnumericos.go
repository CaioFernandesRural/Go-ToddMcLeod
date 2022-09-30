package main

import (
	"fmt"
	"runtime"
)

var x int = 42
var y float64 = 42.12433

func main() {
	//x := 42
	//y := 42.12433

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
