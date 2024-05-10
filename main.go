package main

import (
	"fmt"

	"github.com/juanmontilva/godesde0/variables"
)

func main() {
	fmt.Println("hola mundo")

	variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(950)
	fmt.Println(estado)
	fmt.Println(texto)

}
