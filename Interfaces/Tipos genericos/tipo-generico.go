package main

import "fmt"

//Aceita qualquer tipo
func generica(interf interface{}) {
	fmt.Println(interf)
}
func main() {
	fmt.Println("Tipo Genérico")

	generica("string")
	generica(1)
	generica(true)

	//Exemplo de multiplas interface é o Println que recebe multiplas interfaces
	fmt.Println(1, 2, 100, false, true, float64(123.3))

	//Mapa de tipos genéricos
	mapa := map[interface{}]interface{} {
		1: "string",
		float64(231): true,
		"string": "string",

	}
	fmt.Println(mapa)

}
