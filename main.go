package main

import (
	e "github.com/cajami/godesde0/ejer_interfaces"
	"github.com/cajami/godesde0/modelos"
)

func main() {

	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)

	// users.AltaUsuario()

	// mapas.MostrarMapas()

	// arreglosslices.Capacidad()

	// files.LeoArchivo()
	// files.SumaTabla()
	// texto := ejercicios.Multiplicar()

	// fmt.Println(texto)
	// iteraciones.Iterar()

	// teclado.IngresoNumeros()

	// numero, texto := ejercicios.ConvertirTextoANuevo("199")

	// fmt.Println(numero)
	// fmt.Println(texto)

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
