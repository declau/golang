package main

import "fmt"

//Forma mais aproximada de uma herança que Golang trabalha
type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança!")

	p1 := pessoa{"Denis", "Claudiano", 38}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "PUCMinas"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.idade)
	fmt.Println(e1.faculdade)
}
