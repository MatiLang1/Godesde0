package arreglos_slices

import (
	"fmt"
)

var tabla [10]int = [10]int{10, 0, 0, 59} //asignamos varios valores al array

// var matriz [2][4]int //creación de una matriz (4 elementos 2 arrays)

func MuestroArreglos() {
	tabla[2] = 10 //sobreescribe el 0 que le pusimos arriba
	tabla[4] = 5

	tabla2 := [10]string{"Pablo", "Juan"} //array creado por asignación
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	} //ciclo para mostrar cada elemento por separado del array "tabla"

	// matriz[0][2] = 15
	// fmt.Println(matriz)

}
