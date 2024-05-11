package ejercicios

import "strconv"

func Ejercicio1(letra string) (int, string) {
	num, err := strconv.Atoi(letra)

	// NIL ES NULO, HAY QUE ESTAR ATENTO AL LENGUAJE
	// PORQUE ES BACKEND

	if err != nil {
		return 0, "hubo un error" + err.Error()
	}

	if num > 100 {
		return num, "es mayor a 100"
	} else {
		return num, "es menor de 100"
	}
}
