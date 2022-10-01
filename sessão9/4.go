package main

import "fmt"

func main() {
	bd := 2002
	for {
		fmt.Println(bd)
		bd++
		if bd > 2022 {
			break
		}
	}
}
