package main

import (
	"fmt"

	"../variables/math"
)

// Isso aqui é uma declaração de variável
var a string

func main() {
	// Isso aqui é uma atribuição de variável
	a = "hello"
	println(a)

	// Podemos fazer a declaração e atribuição ao mesmo tempo, dessa forma
	b := "world"
	println(b)

	// Tipos de variáveis
	c := 10
	d := 3.144
	e := true
	f := `
		uoou 

		que legal :D
	`

	fmt.Printf("%T %T %T %T %T\n", c, d, e, f, a)

	resultado := math.Soma(1, 2)

	resultado = math.Multiplicar(2, 2)
	print(resultado)

	resultado2 := math.BomDia
	println(resultado2)
}
