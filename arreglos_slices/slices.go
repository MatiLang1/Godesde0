package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 5, 7, 4, 12, 4, 6, 4, 1}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   //Slice creado desde un vector (desde la posición 3)
	porcion2 := arreglo[:5]  //hasta la posición 5 sin incluir, no incluye el último [)
	porcion3 := arreglo[2:6] //desde la posición 2 al 6
	porcion4 := arreglo[:]

	fmt.Println(porcion, porcion2, porcion3, porcion4)
}

func Capacidad() {
	elementos := make([]int, 5, 20) //tipo,dimensión (cantidad de elementos), capacidad
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ { //se agregan 100 elementos al slice
		nums = append(nums, i) //agrega el elemento "i"
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
