package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumero() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ingrese numero 1 : ")
	if scanner.Scan() {

		// OJO CON ESTE DETALLE
		// LA VARIABLE DA ERROR POR TENER EL :
		// PORQUE LO DA, SIMPLE, PORQUE LA VARIABLE ESTA CREADA
		// POR ESO SE QUITA LOS :
		// numero1, err := strconv.Atoi(scanner.Text())

		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("EL DATO INGRESADO ES INCORRECTO " + err.Error())
		}
	}

	fmt.Println("ingrese numero 2 : ")
	if scanner.Scan() {

		// OJO CON ESTE DETALLE
		// LA VARIABLE DA ERROR POR TENER EL :
		// PORQUE LO DA, SIMPLE, PORQUE LA VARIABLE ESTA CREADA
		// POR ESO SE QUITA LOS :
		// numero1, err := strconv.Atoi(scanner.Text())

		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("EL DATO INGRESADO ES INCORRECTO " + err.Error())
		}
	}

	fmt.Println("ingrese leyenda : ")
	if scanner.Scan() {

		// OJO CON ESTE DETALLE
		// LA VARIABLE DA ERROR POR TENER EL :
		// PORQUE LO DA, SIMPLE, PORQUE LA VARIABLE ESTA CREADA
		// POR ESO SE QUITA LOS :
		// numero1, err := strconv.Atoi(scanner.Text())

		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)

}
