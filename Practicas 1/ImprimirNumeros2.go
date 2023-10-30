package main

import "fmt"

func imprimirRecursivo(numero int) {
	if numero > 0 {
		fmt.Println(numero)
		numero--
		imprimirRecursivo(numero)
	} else {
		fmt.Println(0)
	}
}

func main() {
	imprimirRecursivo(100)
}
