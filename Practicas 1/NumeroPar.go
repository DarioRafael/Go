package main

import (
	"fmt"
)

func main() {

	var numero int

	fmt.Print("Dame el numero:")
	fmt.Scanln(&numero)

	if numero%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}

}
