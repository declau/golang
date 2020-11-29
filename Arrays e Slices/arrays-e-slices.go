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
}
