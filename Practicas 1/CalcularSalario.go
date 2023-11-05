package main

import "fmt"

type Persona struct {
	Salario      float64
	Prestaciones float64
}

func main() {

	persona1 := Persona{Salario: 2, Prestaciones: 20}
	persona2 := Persona{Salario: 2, Prestaciones: 20}

	fmt.Println(persona1)
	fmt.Println(persona2)

}
