package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string, 5)
	// fmt.Println(paises)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	paises["Venezuela"] = "Caracas"

	fmt.Println(paises)
	fmt.Println(paises["Venezuela"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("equipo %s, tiene un puntaje de %d \n \n", equipo, puntaje)
	}

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t ", puntaje, existe)

}
