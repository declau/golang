package main

import "fmt"

func main() {
	fmt.Println("If Else")

	//number := 10
	number := -5
	//number := -11

	if number > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//if init
	//Ao criar um variável pelo if init ela fica visível somente dentro do escopo do if
	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("Número é maior que 0")
	} else if anotherNumber < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número esta entre 0 e -10")
	}
}
