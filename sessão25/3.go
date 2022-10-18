package main

import (
	"fmt"
)

type customErr struct {
	erro string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("ó o erro %v", ce.erro)
}

func main() {
	c1 := customErr{"tô sem saco"}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
