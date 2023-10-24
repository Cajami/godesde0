package arreglosslices

import "fmt"

var tablaS []int = []int{2, 5, 6}

var arreglo [10]int = [10]int{78, 6, 5, 8, 4, 1, 155, 5, 7, 8}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:] // Slice creado desde un vector desde la posición 3

	porcion2 := arreglo[:5] //Slice creado desde la posición cero hasta la 5

	porcion3 := arreglo[6:7] //Slice creado desde un vefctor desde la posicion 6 hasta las 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad(){
	elementos := make([]int, 5, 20)

	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make ([]int, 0,0)

	for i:=0;i<100;i++{
		nums = append(nums,i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
