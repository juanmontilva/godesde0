package ejercicios

import (
	"strconv"
)

func Ejercicio1(letra string) (int, string) {
	num, err := strconv.Atoi(letra)

	// NIL ES NULO, HAY QUE ESTAR ATENTO AL LENGUAJE
	// PORQUE ES BACKEND

	if err != nil {
		return 0, "hubo un error " + err.Error()
	}

	// AMBAS FORMAS DE DECLARAR FUNCIONAN

	switch {
	case num > 100:
		return num, "es mayor a 100"
	case num < 100:
		return num, "es menor a 100"
	default:
		return num, "es igual a 100"
	}

	// if num > 100 {
	// 	return num, "es mayor a 100"
	// } else {
	// 	return num, "es menor de 100"
	// }

}
