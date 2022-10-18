package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	log.SetOutput(f)

	f1, err := os.Open("a.txt")

	if err != nil {
		log.Println("teve erro ", err)
	}
	defer f1.Close()
}
