package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

//Função com ponteiro
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Funções com Ponteiros")
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	fmt.Println("---------------------")

	//resultado usando função com ponteiro
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
