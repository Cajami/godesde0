package main

import (
	// deferpanic "github.com/cajami/godesde0/defer_panic"

	"fmt"

	"github.com/cajami/godesde0/goroutines"
)

func main() {

	canal1 := make(chan bool)

	go goroutines.MiNombreLento("Javier", canal1)

	defer func() {
		<-canal1

	}()

	fmt.Println("Estoy aqui")

	// var x string
	// fmt.Scanln(&x)

	// deferpanic.VemosDefer()
	// deferpanic.EjemploPanic()

	// Pedro := new(modelos.Hombre)
	// e.HumanosRespirando(Pedro)

	// Maria := new(modelos.Mujer)
	// e.HumanosRespirando(Maria)

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
