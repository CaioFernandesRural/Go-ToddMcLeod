package main

import "fmt"

func main() {
	jb := []string{"banana", "morango", "chocolate", "caramelo"}

	fmt.Println(jb)

	mp := []string{"baunilha", "creme", "menta", "ovomaltine"}

	fmt.Println(mp)

	xp := [][]string{jb, mp}

	fmt.Println(xp)
}
