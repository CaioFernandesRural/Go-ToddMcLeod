package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("arquivo.txt")

	if err != nil {
		//fmt.Println("houve erro", err)
		//log.Println(err)
		log.Fatalln(err)
		//panic(err)
	}
}
