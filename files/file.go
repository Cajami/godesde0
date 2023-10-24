package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cajami/godesde0/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Multiplicar()
	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo ", err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Multiplicar()

	if Append(texto) == false {
		fmt.Println("Error al concatenar contenido")
	} else {
		fmt.Println("Se grabaron los datos al final del archivo")

	}

}

func Append(texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante al Append ", err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el WriteString ", err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error al leer el archivo ", err.Error())
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(registro)

	}
	archivo.Close()

}
