package iteraciones

import "fmt"

func Iterar() {

	// VARIOS EJEMPLOS DEL CICLO FOR,
	// OJO EN GOLANG NO SE USA WHILE,
	// SOLAMENTE FOR Y FUNCIONA PARA TODO

	// for {
	// 	fmt.Println("hola desde el ciclo for")
	// 	break
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 100; i += 10 {
	// 	fmt.Println(i)
	// }

	// for i := 100; i > 10; i -= 10 {
	// 	fmt.Println(i)
	// }

	// for i := 10; i > 1; i-- {
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	for i := 10; i > 1; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}

}
