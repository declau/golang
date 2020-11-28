package main

import "fmt"

func main() {

	//Aritimeticos
	fmt.Println("Aritimeticos---------------------------------")
	soma := 2 + 1
	subtracao := 2 - 1
	multiplicacao := 2 * 2
	divisao := 10 / 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 15
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	//fim aritimeticos

	//Atribuição
	fmt.Println("Atribuição---------------------------------")
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//Fim dos operadores atribuição

	//Operadores relacionais
	//sempre retorna true ou false
	fmt.Println("relacionais---------------------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	//Fim dos relacionais

	//Operadores Lógicos
	fmt.Println("Lógicos---------------------------------")
	verdadeirao, falso := true, false
	fmt.Println(verdadeirao && falso) //verifica se todas as condições são true
	fmt.Println(verdadeirao || falso) //verifica se uma das condições são true
	fmt.Println(!verdadeirao)         //negação
	fmt.Println(!falso)               //negação
	//Fim dos operadores lógicos

	//Operadores Unários
	fmt.Println("Unários---------------------------------")
	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)
	numero--
	numero -= 10 //numero = numero + 10
	fmt.Println(numero)
	numero *= 3 //numero = numero * 3
	fmt.Println(numero)
	numero /= 10 //numero = numero / 10
	fmt.Println(numero)
	numero %= 3 //numero = numero %= 3

}
