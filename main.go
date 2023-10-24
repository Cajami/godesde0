package main

import (
	"fmt"

	"github.com/cajami/godesde0/ejercicios"
)

func main() {

	numero, texto := ejercicios.ConvertirTextoANuevo("199")

	fmt.Println(numero)
	fmt.Println(texto)

	// estado, texto := variables.ConviertoaTexto(216)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// os := runtime.GOOS

	// if os == "Linux." {
	// 	fmt.Println("Es un linux")
	// } else {
	// 	fmt.Print("El sistema encontrado es: ", os)
	// }

}
