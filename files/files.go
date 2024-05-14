package files

import (
	"bufio"
	"fmt"

	// OBSOLETO, AHORA SE ENCARGA DE ESTO OS
	// "io/ioutil"

	"os"

	"github.com/juanmontilva/godesde0/ejercicios"
)

var RutaArchivo string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.PedirNumero()
	archivo, err := os.Create(RutaArchivo)
	if err != nil {
		fmt.Println("error al mostrar el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	// OJO TODO ARCHIVO PARA LEER O GRABAR TIENE QUE CERRARSE
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.PedirNumero()
	// ESTO ME PARECE EXCELENTE, EN LA CLASE VI QUE TENIA QUE PONERLO FALSO, CON LA SINTAXIS DE AYUDA DE GO, NO ME HACIA FALTA COLOCARLE FALSO, SOLAMENTE SI DIFERENCIA, EJEMPLO !, NO HAY QUE PONERLE FALSE
	// if Append(RutaArchivo, texto) == false {
	if !Append(RutaArchivo, texto) {
		fmt.Println("error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(RutaArchivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("erro durante el append" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("erro durante el writeString" + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	// OJO ESTA OBSOLETO, AHORA SE ENCARGA OS DEL ASUNTO

	// archivo, err := os.ReadFile(RutaArchivo)
	// if err != nil {
	// 	fmt.Println("error al leer archivo" + err.Error())
	// 	return
	// }

	// fmt.Println(string(archivo))

	// OTRA ALTERNATIVA
	archivo, err := os.Open(RutaArchivo)
	if err != nil {
		fmt.Println("error al leer el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	archivo.Close()

}
