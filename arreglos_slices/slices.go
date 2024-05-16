package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 99, 2, 5, 67, 7, 8, 3, 1, 0}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //Slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  //slice creado desde un vector, desde la posicion 0 hasta la5
	porcion3 := arreglo[6:7] //slice creado desde un vector, desde la posicion 6 hata la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}
