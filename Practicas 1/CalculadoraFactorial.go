package main

import "fmt"

func main() {
	num := 0
	resultado := 1

	fmt.Print("Dame el numero entero...")
	fmt.Scanln(&num)

	for i := num; i >= 1; i-- {
		resultado *= i
		fmt.Println(i, resultado)
	}

}
