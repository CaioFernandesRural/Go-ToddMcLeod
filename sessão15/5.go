package main

import (
	"fmt"
	"math"
)

type square struct {
	lado float64
}
type circle struct {
	raio float64
}

func (s square) area() float64 {
	return s.lado * s.lado
}
func (c circle) area() float64 {
	return c.raio * c.raio * math.Pi
}

type shape interface {
	area() float64
}

func info(x shape) {
	fmt.Println(x.area())
}

func main() {
	c := circle{
		raio: 12.5,
	}
	s := square{
		lado: 7.3,
	}
	info(c)
	info(s)

}
