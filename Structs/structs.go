package main

import "fmt"

//criando um tipo personalizado Usuário
//Structs seria como se fosse a classe em OO
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo-Structs")

	var user1 usuario
	user1.nome = "Denis"
	user1.idade = 38
	fmt.Println("Nome:", user1.nome, "Idade:", user1.idade)

	enderecoExemplo := endereco{"rua Chico Balua", 45}

	user2 := usuario{"Eduardo", 30, enderecoExemplo}
	fmt.Println("Nome:", user2.nome, "Idade:", user2.idade, "Endereço", user2.endereco)
	fmt.Println("struct dentro de struct", user2)

	user3 := usuario{nome: "Declau"}
	fmt.Println("Nome:", user3.nome, "Idade:", user3.idade)
}
