package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("a.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	io.Copy(f, r)
}
