package main

import "fmt"

func main() {
	m := map[string][]string{
		"bonde_james": {"martini", "carros", "wemen"},
		"claudin":     {"a", "b", "c"},
		"L":           {"13", "vermelho", "gozo"},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		for i, va := range v {
			fmt.Println(i, va)
		}
	}

	m["caio"] = []string{"dinheiro", "musica", "hardware"}
	fmt.Println(m)
}
