package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicar() {
	var numero int
	var err error

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

	fmt.Println("Tabla Multiplicar del ", numero)
	fmt.Println("")

	for i := 1; i <= 10; i++ {
		fmt.Println(numero, " x ", i, " = ", numero*i)
	}

}
