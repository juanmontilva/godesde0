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
