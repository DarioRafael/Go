package main

import "fmt"

func numero(a int) {
	if a <= 100 {
		// Imprime el número actual
		fmt.Println(a)
		// Llamada recursiva para el siguiente número
		numero(a + 1)
	}
}

func main() {

	fmt.Println("Prueba 1.")

	fmt.Println("Forma iterativa")
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n\n")
	fmt.Println("Forma recursiva")
	numero(1)

}
