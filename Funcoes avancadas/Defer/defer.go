package main

import (
	"fmt"
)

//A instrução defer adiciona a chamada da função após a palavra-chave defer em uma pilha

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. O resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Função Defer")

	//Defer = Adiar
	/* defer funcao1()
	funcao2() */

	fmt.Println("---------------------------")
	fmt.Println("Resultado Aluno: ", alunoEstaAprovado(7, 8))

	//Outro exemplo simples
	//fmt.Println("---------------------------")
	//defer fmt.Println("Bye")
	//fmt.Println("Hi")
}
