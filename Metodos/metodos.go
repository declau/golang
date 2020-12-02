package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//Método utilizando ponteiro
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")
	fmt.Println("---------------------")

	usuario1 := usuario{"Denis", 38}
	//fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Eduardo", 22}
	usuario2.salvar()
	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	//Fazendo aniversário
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
