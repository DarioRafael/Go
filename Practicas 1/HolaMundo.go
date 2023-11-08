package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myString = "Hola Mundo"
	var myInt = 2
	var myFloat = 2.6

	fmt.Println(myString, myInt, myFloat)

	fmt.Println(myFloat + float64(myInt))

	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println("Soy el mejor")
}
