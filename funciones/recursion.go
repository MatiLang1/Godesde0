package funciones

import "fmt"

func Exponencial(valor int) {
	if valor > 10000000 {
		return //cuando valor supera a "10000000" sale de la función
	}
	fmt.Println(valor)
	Exponencial(valor * 2) //se multiplica por 2 y se muestra el valor por pantalla hasta llegar al valor "10000000"
}
