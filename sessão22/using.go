package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	bar(c)

	fmt.Println("boutta head out")
}
func foo(c chan<- int) { //manda
	c <- 42
}
func bar(c <-chan int) { //recebe
	fmt.Println(<-c)
}
