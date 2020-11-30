package main

import (
	"fmt"
)

//Função de retorno nomeado
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println("Retorno Nomeado")
	soma, subtracao := calculosMatematicos(10, 5)
	fmt.Println(soma, subtracao)

}
