package main

import "fmt"

//Não podemos ter mais de um parâmetro variático por função e
//obrigatoriamente tem que ser o último parâmetro
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}

//Não podemos ter mais de um parâmetro variático por função e
//obrigatoriamente tem que ser o último parâmetro
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções Variáticas")
	totalSoma := soma(1, 2, 3, 4, 5, 6)
	fmt.Println(totalSoma)

	escrever("Olá", 1, 3, 5)
}
