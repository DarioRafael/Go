package main

import "fmt"

type Rectangulo struct {
	Ancho  float64
	Altura float64
}

func (rec Rectangulo) CalcularArea() float64 {
	return rec.Ancho * rec.Altura
}

/*func (receptor TipoDeDatos) NombreDelMetodo() TipoDeRetorno {
	// Cuerpo del método
}*/


func main() {

	rect := Rectangulo{Ancho: 10, Altura: 5}
	area := rect.CalcularArea()
	fmt.Println("Área del rectángulo:", area)

}
