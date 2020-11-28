package main

import (
	"fmt"
)

//Funções são tipos em Golang
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//retornos multiplos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var d = func(txt string) {
		fmt.Println(txt)
	}
	d("Texto da função d")

	var e = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := e("Texto da função e")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//omitindo o resultado da subtração
	resultadoSoma2, _ := calculosMatematicos(10, 12)
	fmt.Println(resultadoSoma2)
}
