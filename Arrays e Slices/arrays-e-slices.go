package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//Arrays
	var array1 [5]string
	array1[0] = "posição 1"
	fmt.Println(array1)

	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//Slices
	slice := []int{10, 11, 3, 9, 10, 11}
	fmt.Println(slice)

	//reflect.TypeOf devolve o tipo de uma variável
	//Slice tem o tamnho dinâmico
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//Usando Append, adiciona-se um item no Slice e retorna um novo slice com o item adicionado
	slice = append(slice, 7)
	fmt.Println(slice)

	//Referenciar um Array
	//Slice2 pega uma fatia do arry2 criando um novo Slice
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	fmt.Println("----------------------------------")
	//ARRAYS INTERNOS
	//Função make aloca um espaço na memória
	//Make recebe três parâmetros(tipo, tamanho, capacidade máxima)
	//Função make inicialmente vai criar um array de 11 posições
	//e vai retornar um Slice que vai pegar as 10 primeiras posições desse Array
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length, mostra o tamanho o Slice
	fmt.Println(cap(slice3)) //capacity, mostra a capacidade do Slice

	//Estourar um Slice
	slice3 = append(slice3, 3)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length, mostra o tamanho o Slice
	fmt.Println(cap(slice3)) //capacity, mostra a capacidade do Slice
	// Tamanho 11 , Capacidade 11
	//Vai estourar????
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length, mostra o tamanho o Slice
	fmt.Println(cap(slice3)) //capacity, mostra a capacidade do Slice
	//Não estourou. Tamanho foi para 12 e capacidade 24.
	//Go quando ve que o Slice vai estourar ele cria um outro Array para se referenciar
	//e dobra seu tamanho

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) //length, mostra o tamanho o Slice
	fmt.Println(cap(slice4)) //capacity, mostra a capacidade do Slice

}
