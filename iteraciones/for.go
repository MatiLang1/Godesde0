package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 10; i++ { //se declara "i" dentro del for, mientras sea menor a 10 se ejecuta, luego i++
		if i == 6 {
			break //aborta el for
		}
		fmt.Println(i)
	}
}

func Iterar2() {
	for i := 10; i > 1; i-- { //se declara "i" dentro del for, mientras sea menor a 10 se ejecuta, luego i++
		if i == 6 {
			continue //vuelve arriba al for a evaluar la instruccion (no sigue con lo debajo reinicia el ciclo for)
		}
		fmt.Println(i)
	}
}
