package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipos de int int8, int16, int32, int64
	var numero int32 = 1000000000
	fmt.Println(numero)

	var numero2 uint64 = 10000000000000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 1000000000
	fmt.Println(numero3)

	//uint32 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	//números reais
	//numeros reais  float32, float64
	var numeroReal1 float32 = 123.12
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234.456789
	fmt.Println(numeroReal2)

	//inferencia de tipo
	numeroReal3 := 1234578.123
	fmt.Println(numeroReal3)
	//fim números reais

	//Strings

	var str string = "blablabla"
	fmt.Println(str)

	str2 := "kkkkkk"
	fmt.Println(str2)

	//char
	char := 'B'
	//Pega o número referente ao caracter da tabela ASCII
	fmt.Println(char)

	//fim strings

	//valor zero
	var texto string
	//retorna valor vazio para strings
	fmt.Println(texto)

	var texto2 int32
	//retorna valor zero para números
	fmt.Println(texto2)

	//para erro retorna o nil
	//fim valor zero

	//boolean
	var booleano bool = true
	fmt.Println(booleano)

	//sem declarar o valor é sempre false
	var booleano2 bool
	fmt.Println(booleano2)
	//fim boolean

	//tipo erro
	var erro error
	fmt.Println(erro)

	//criar um erro personalizado
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
