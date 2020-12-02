package main

import (
	"fmt"
	"math"
)

//Interface
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("Á área da forma é %0.2f \n", f.area())
}

//Médoto que calcula área do retangulo
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

//Método que calcula área do circulo
func (c circulo) area() float64 {
	//return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2)
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func main() {
	fmt.Println("Interfaces")
	r := retangulo{10, 10}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
