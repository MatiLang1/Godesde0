package funciones

import "fmt"

func Calculos() {

	var n3 int = 32
	var n4 int = 240

	calculo := func(n1, n2 int) int {
		//asignamos una función a una variable
		return n1 + n2 + n3 + n4
		//si la función "suma" se definiera fuera de "Calculos", ésta no tendría acceso a las variables n3 y n4,
		// ya que estas solo tienen scope (alcance) dentro de la función "Calculos"
	}

	fmt.Println(calculo(10, 20))

	calculo = func(n1 int, n2 int) int {
		// reutilizamos la misma funcion para realizar otra cosa (tiene otro return) pero debe utilizar los mismos argumentos
		return n1 * n2 * n3
	}

	fmt.Println(calculo(10, 7))
	//le pasamos los parametros a la funcion y mostramos por pantalla

}
