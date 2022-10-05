package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}
type picape struct {
	veiculo
	quatrop4 bool
}
type seda struct {
	veiculo
	luxo bool
}

func main() {
	a := picape{
		veiculo: veiculo{
			portas: 2,
			cor:    "preto",
		},
		quatrop4: true,
	}
	b := seda{
		veiculo: veiculo{
			portas: 4,
			cor:    "vermelho",
		},
		luxo: false,
	}
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a.portas, a.cor, a.quatrop4)
	fmt.Println(b.portas, b.cor, b.luxo)
}
