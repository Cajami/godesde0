package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicar() string {
	var numero int
	var err error
	var texto string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un número: ")

		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("El dato ingresado es incorrecto")
				continue
			}
			break
		}

	}

	texto += fmt.Sprintln("-------------------------------------")
	texto += fmt.Sprintln("Tabla Multiplicar del ", numero)
	texto += fmt.Sprintln("")

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintln(numero, " x ", i, " = ", numero*i)
	}
	texto += fmt.Sprintln("-------------------------------------")

	return texto
}
