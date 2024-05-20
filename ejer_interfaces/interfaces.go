package ejer_interfaces

import (
	"fmt"

	"github.com/juanmontilva/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	hu.EstaVivo()
	fmt.Printf("soy un/a %s y estoy respirando, trambie estoy vivo = %t \n", hu.Sexo(), hu.EstaVivo())
}

func TipoEspecie(Ani interfaces.Animal) {
	Ani.Especie()
	Ani.EstaVivo()
	Ani.Comer()
	Ani.EsCarnivoro()
	fmt.Printf("soy un %s y soy %t carnivoro y como %s y estoy vivo %t\n", Ani.Especie(), Ani.EsCarnivoro(), Ani.Comer(), Ani.EstaVivo())
}

func TipoVegetal(Vegeta interfaces.Vegetal) {
	Vegeta.ClasificacionVegetal()
	Vegeta.EstaVivo()
	fmt.Printf("soy un vegetal tipo %s y estoy vivo %t \n", Vegeta.ClasificacionVegetal(), Vegeta.EstaVivo())
}
