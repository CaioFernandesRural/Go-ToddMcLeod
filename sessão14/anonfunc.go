package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("função anônima")
	}()
	func(x int) {
		fmt.Println("o sentido da vida, ", x)
	}(42)
}

func foo() {
	fmt.Println("foo")
}
