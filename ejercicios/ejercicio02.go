package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TabladeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin) //prepara el escáner para recibir lo que el usuario escriba

	for {
		fmt.Println("Ingrese un número: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text()) //convierte de string a int
			if err != nil {
				continue //hay un error en el n ingresado, se lo vuelve a pedir
			} else {
				break //ya pidió el numero sale del ciclo
			}
		}
	}

	for i := 1; i < 11; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero) //se almacena en la variable "texto" toda la tabla que se genera (asi se tiene un archivo almacenado en una variable y se puede llamar a la tabla entera usando la variable)
	}

	return texto
}
