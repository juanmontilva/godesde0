package main

import (
	"fmt"
	"runtime"

	"github.com/juanmontilva/godesde0/ejercicios"
	"github.com/juanmontilva/godesde0/variables"
)

// %s: para cadenas de texto.
// %d: para enteros.
// %f: para números de punto flotante.
// %v: para el valor por defecto del tipo de dato.
// %t: para valores booleanos.
// %e y %E: para números de punto flotante en formato científico.
// Por otro lado, \n es un carácter de escape que representa una nueva línea. Se utiliza para insertar un salto de línea en una cadena de texto.

func main() {
	fmt.Println("hola mundo")

	variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(950)
	fmt.Println(estado)
	fmt.Println(texto)

	// -------------------------------
	if os := runtime.GOOS; os == "Linux" || os == "OSX" || os == "darwin" {
		fmt.Println("1 Esto es un = ", os)
	} else {
		fmt.Println("es windows o mac", os)
	}

	// ----------------------------
	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("esto es = linux")
	case "windows":
		fmt.Println("esto es = windows")
	default:
		fmt.Printf("esto es otro software y es %s \n ", os)
	}

	// --------------------------

	numero, texto := ejercicios.Ejercicio1("1f0")
	fmt.Println(numero)
	fmt.Println(texto)
}
