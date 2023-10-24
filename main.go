package main

import (
	"fmt"

	"github.com/cajami/godesde0/variables"
)

func main() {

	estado, texto := variables.ConviertoaTexto(216)
	fmt.Println(estado)
	fmt.Println(texto)

}
