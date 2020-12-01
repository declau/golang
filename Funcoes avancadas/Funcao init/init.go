package main

import "fmt"

var n int

//É simplismente uma função que vai ser executada antes da função Main
func init() {
	fmt.Println("Executando a função Init....")
	n = 10
}

func main() {
	fmt.Println("Função Init")
	fmt.Println("Executando função main....")
	fmt.Println(n)
}
