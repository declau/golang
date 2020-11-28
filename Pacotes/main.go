package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("alguma coisas do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("dec@gmail.com")
	fmt.Println(erro)
}
