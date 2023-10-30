package main

import (
	"fmt"
)

type Triangulo struct {
	Base   float64
	Altura float64
}

func (tri Triangulo) calcularArea() float64 {
	return (tri.Base * tri.Altura) / 2
}

func main() {
	var altura float64
	var base float64

	fmt.Print("Altura: ")
	fmt.Scanln(&altura)
	fmt.Print("Base: ")
	fmt.Scanln(&base)

	tri := Triangulo{Base: base, Altura: altura}

	area := tri.calcularArea()
	fmt.Print("Area de un triángulo ", area)

}
