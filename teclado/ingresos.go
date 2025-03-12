package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//todo lo que entra por bufio es de tipo string por lo que hay que convertir

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingrese número 1: ")
	scanner := bufio.NewScanner(os.Stdin) //stdin es standard input (el teclado)
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text()) //convertimos string a int

		if err != nil {
			panic("El dato ingresado es incorrecto") // sale de la aplicación porque ocurrió un error
		}
	}

	fmt.Println("Ingrese número 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado es incorrecto")
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)

}
