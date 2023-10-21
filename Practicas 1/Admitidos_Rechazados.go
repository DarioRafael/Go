package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString string
	var myAge int
	var myDni string

	admitidos := []string{"DARIO", "MARIO"}
	rechazados := []string{"MARTHA", "PEDRO"}

	fmt.Print("Nombre...")
	fmt.Scanln(&myString)

	if myString == "" {
		fmt.Println("Nombre no valido... Cerrando programa")
		return
	}

	fmt.Print("Edad...")
	fmt.Scanln(&myAge)

	if !(myAge > 0 && myAge < 120) {
		fmt.Println("Edad no aceptada... Cerrando programa")
		return
	}

	fmt.Print("Dni...")
	fmt.Scanln(&myDni)

	if !(len(myDni) == 10) {
		fmt.Println("DNI no aceptado... Cerrando programa")
		return
	}

	fmt.Println("\nNombre:"+myString+
		"\nEdad:", myAge,
		"\nMatricula:", myDni)

	//NombreVerificar := strings.ToUpper(myString)
	var NombreVerificar string = strings.ToUpper(myString)
	var enProceso bool = true

	for i := 0; i < len(admitidos) && i < len(rechazados); i++ {
		if admitidos[i] == NombreVerificar {
			fmt.Println("ADMITIDO")
			enProceso = false
		} else if rechazados[i] == NombreVerificar {
			fmt.Println("RECHAZADO")
			enProceso = false
		}
	}

	if enProceso {
		fmt.Println("EN PROCESO DE ADMISION")
	}
}
