package variables

import "fmt"

// hay dos tipos de variables
// declarativa y por asignacion directa

func MuestroEnteros() {
	// variable por declaracion
	var intcomun int

	// por asignacion directa
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)

	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)

}
