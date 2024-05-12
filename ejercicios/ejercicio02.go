package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numeroPrincipal int
var err error

func PedirNumero() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ingrese el numero para multiplicar : ")

	for {
		if scanner.Scan() {
			numeroPrincipal, err = strconv.Atoi(scanner.Text())
			if err != nil {
				// panic("DATO INGRESADO INCORRECTAMENTE" + err.Error())
				continue
			} else {
				break
			}
		}

	}

	for i := 0; i <= 10; i++ {
		multiplicacion := numeroPrincipal * i
		fmt.Printf("%d x %d = %d\n", numeroPrincipal, i, multiplicacion)
	}

	// OJO TAMBIEN FUNCIONA DE ESTA FORMA, ES MAS CORTO
	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("%d x %d = %d \n", numeroPrincipal, i, numeroPrincipal*i)
	// }

}
