package main

import "fmt"

func main() {

	favSport := "strongman"

	switch favSport {
	case "crossfit":
		fmt.Println("frango")
	case "polo":
		fmt.Println("branco")
	case "strongman":
		fmt.Println("brabo")
	default:
		fmt.Println("frango")
	}
}
