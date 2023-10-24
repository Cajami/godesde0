package ejercicios

import (
	"strconv"
)

func ConvertirTextoANuevo(texto string) (int, string) {
	numero, _ := strconv.Atoi(texto)
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
