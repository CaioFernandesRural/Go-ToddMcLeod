package main

import "fmt"

func main() {
	foo()
	bar("roger")
	s1 := woo("Joel")
	fmt.Println(s1)
	x, y := mouse("Ian", "Flemming")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identifier (parâmetros) (return/s) { código }

func foo() {
	fmt.Println("Olá, foo")
}

//tudo no go é passado por valor

func bar(s string) {
	fmt.Println("olá,", s)
}

func woo(s string) string {
	return fmt.Sprint("Olá, ", s)
}
func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, ", says hello")
	b := true
	return a, b
}
