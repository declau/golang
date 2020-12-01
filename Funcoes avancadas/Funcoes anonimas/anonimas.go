package main

import "fmt"

func main() {
	fmt.Println("Funções Anônimas")
	fmt.Println("---------------------------------")

	//Declara e executa a função anônima
	func() {
		fmt.Println("Olá, estou aqui....")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando valor do parâmetro")

	///retornando um string
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido... -> %s", texto)
	}("Passando valor do parâmetro")

	fmt.Println(retorno)
}
