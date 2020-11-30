package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//dentro do colchetes é o tipo das chaves
	//e fora do colchetes é o tipo dos valores
	user := map[string]string{
		"nome":      "Denis",
		"sobrenome": "Claudiano",
	}
	fmt.Println(user)

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(user2)

	//Deletar uma chave o Map
	delete(user2, "nome")
	fmt.Println(user2)

	//Adicionar uma chave
	user2["signo"] = map[string]string{
		"nome": "Escorpião",
	}
	fmt.Println(user2["signo"])
}
