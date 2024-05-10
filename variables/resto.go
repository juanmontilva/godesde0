// OJO, LAS VARIABLES SE TIENEN QUE DECLARAR EN MAYUSCULA
// AL PRINCIPIO PARA QUE SEA TIPO PUBLICA

package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "pedro"
	Estado = true
	Sueldo = 1577.89
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)

}

// FUNCIONES

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
