package ejercicios

import (
	"strconv"
)

func Ejercicio1(valor string) (int, string) {
	numero, err := strconv.Atoi(valor) // Captura tanto el número como el error
	if err != nil {                    // Maneja el error en caso de que el valor no sea un número válido
		return 0, "Error: el valor ingresado no es un número válido" + err.Error() //se retorna el error en tipo string, con el + se concatena
		//siempre se retorna un numero y un texto (int,string ya que eso retorna nuestra función)

	}

	if numero > 100 {
		valorstring := "El valor es mayor a 100"
		return numero, valorstring
		//return numero, "El valor es mayor a 100" (esta opción es mejor)

	} else {
		valorstring := "El valor no es mayor a 100"
		return numero, valorstring
	}

}
