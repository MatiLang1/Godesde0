package ejercicios

import (
	"strconv"
)

func Ejercicio1(valor string) (int, string) {
	numero, err := strconv.Atoi(valor) // Captura tanto el número como el error
	if err != nil {                    // Maneja el error en caso de que el valor no sea un número válido
		return 0, "Error: el valor ingresado no es un número válido" //siempre se retorna un numero y un texto (int,string ya que eso retorna nuestra función)
	}

	if numero > 100 {
		valorstring := "El valor es mayor a 100"
		return numero, valorstring
	} else {
		valorstring := "El valor no es mayor a 100"
		return numero, valorstring
	}

}
