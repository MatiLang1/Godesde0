package variables

import (
	"fmt"
	"strconv"
	"time"
)

//var Nombre string tiene scope global en todo el paquete
//var nombre scope global solo en este archivo

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	//var nombre scope local dentro de ésta función
	// MuestroEnteros() funcion de otro archivo dentro del mismo package
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1988.99
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)

}

func ConviertoaTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto //devuelve true y el texto
}
