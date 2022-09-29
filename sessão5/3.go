package main

import "fmt"

var x int = 42
var y string = "James bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v %v %v\n", x, y, z)
	fmt.Println(s)
}
