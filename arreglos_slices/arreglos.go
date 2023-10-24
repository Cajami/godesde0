package arreglosslices

import "fmt"

var tabla [10]int = [10]int{10, 40, 59}
var matriz [30][30]int

func MuestroArreglos() {
	tabla[7] = 65
	tabla[4] = 45

	tabla2 := [10]string{"Javier", "Fabrizio"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[5][12] = 15

	fmt.Println(matriz)

}
