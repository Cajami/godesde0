package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	// fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenso Aires"

	// fmt.Println(paises)
	// fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona" : 15,
		"Real Madrid":74,
		"Chivas":91,
		"Boca Juniors":65,
	}

	fmt.Println(campeonato)

	// for equipo, puntaje :=range campeonato{
	// 	fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	// }

	delete(campeonato, "Real Madrid")

	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]

	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t", puntaje, existe)


}
