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

	/*for i := 0; i <= 100; i++ {

		fmt.Println(i)
	}*/

	numero(1)

}
