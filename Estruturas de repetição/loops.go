package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Estruturas de Repetição")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i ...", i)
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	//Outra sintaxe
	//Aqui o j vai ficar visível somente dentro do escopo do For
	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j ...", j)
		time.Sleep(time.Second)
	}

	//For com cláusola Range, usado quando for preciso iterar por exemplo em um Array, Slice ...
	nomes := [3]string{"Denis", "Eduardo", "Claudiano"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("-----------------------")

	//Exemplo com Slice
	nomes2 := []string{"Denis", "Eduardo", "Claudiano"}
	//Omitindo o indice
	for _, nome := range nomes2 {
		fmt.Println(nome)
	}

	fmt.Println("-----------------------")

	//Iterando por uma String
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	fmt.Println("-----------------------")

	//Iterando por um Map
	//Não é possível utilizar o Range em um Structs
	usuario := map[string]string{
		"nome":      "Marcus",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
