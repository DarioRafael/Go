package main

import "fmt"

func main() {
	var num, num2, opcion int

	fmt.Print("Primer digito...")
	fmt.Scanln(&num)

	fmt.Print("Segundo digito...")
	fmt.Scanln(&num2)

	fmt.Println("1- Suma")
	fmt.Println("2- Resta")
	fmt.Println("3- Multiplicación")
	fmt.Println("4- División")
	fmt.Println("5- Exponente")

	fmt.Print("Selecciona la opción que gustes...")
	fmt.Scanln(&opcion)

	if opcion == 1 {
		fmt.Println(num + num2)
	} else if opcion == 3 {
		fmt.Println(num - num2)
	} else if opcion == 4 {
		fmt.Println(num * num2)
	} else if opcion == 5 {
		fmt.Println(num / num2)
	} else if opcion == 6 {
		fmt.Println(num)
	}

}
