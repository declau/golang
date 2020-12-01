package main

import "fmt"

func closure() func() {
	texto := "Dentro da função Closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	fmt.Println("Função Closure")
	texto := "Dentro da função Main"
	fmt.Println(texto)

	novaFuncao := closure()
	novaFuncao()
}
