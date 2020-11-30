package main

import (
	"fmt"
)

//Função que vai retornar o dia da semana
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Número Inválido"
	}
}

//Outra forma de Switch
//Essa nomenclatura é mais útil quando não se quer avaliar a mesma variável
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Número Inválido"
	}
}

//Fechando com fallthrough
//fallthrough passa seu código para próxima condição
func diaDaSemana3(numero int) string {

	var diaDaSemana string

	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough
	case 2:
		diaDaSemana = "Segunda-Feira" //setando o valor 1, com fallthrough para aqui :)
	case 3:
		diaDaSemana = "Terça-Feira"
	case 4:
		diaDaSemana = "Quarta-Feira"
	case 5:
		diaDaSemana = "Quinta-Feira"
	case 6:
		diaDaSemana = "Sexta-Feira"
	case 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Número Inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(7)
	fmt.Println(dia2)

	dia3 := diaDaSemana3(5)
	fmt.Println(dia3)
}
