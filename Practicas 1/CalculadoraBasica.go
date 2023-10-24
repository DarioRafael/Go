package main

import "fmt"

func pedirNumeros() (int, int) {
	var num1, num2 int
	fmt.Print("Primer digito: ")
	fmt.Scanln(&num1)
	fmt.Print("Segundo digito: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func main() {
	var opcion int

	fmt.Println("1- Suma")
	fmt.Println("2- Resta")
	fmt.Println("3- Multiplicación")
	fmt.Println("4- División")
	fmt.Println("5- Exponente")

	fmt.Print("Selecciona la opción que gustes...")
	fmt.Scanln(&opcion)

	num, num2 := pedirNumeros()
	if opcion == 1 {
		fmt.Println(num + num2)
	} else if opcion == 2 {
		fmt.Println(num - num2)
	} else if opcion == 3 {
		fmt.Println(num * num2)
	} else if opcion == 4 {
		fmt.Println(num / num2)
	} else if opcion == 5 {
		resultado := 1
		for i := 0; i < num2; i++ {
			resultado *= num
		}
		fmt.Println(resultado)
	}

}
