package main

import "fmt"

func formaRecursiva(num, resultado int) {
	if num < 1 {
		fmt.Println("Factorial:", resultado)
		return
	}
	resultado *= num
	fmt.Println(num, resultado)
	formaRecursiva(num-1, resultado)
}

func main() {
	/*num := 0
	resultado := 1

	fmt.Print("Dame el numero entero...")
	fmt.Scanln(&num)

	for i := num; i >= 1; i-- {
		resultado *= i
		fmt.Println(i, resultado)
	}
	*/
	fmt.Println("\n")
	formaRecursiva(5, 1)

}
