package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

// func TabladeMultiplicar() {
// 	fmt.Println("Ingresar un número: ")
// 	scanner := bufio.NewScanner(os.Stdin) //stdin es standard input (el teclado)
// 	if scanner.Scan() {
// 		numero, err = strconv.Atoi(scanner.Text()) //convertimos string a int

// 		if err != nil {
// 			fmt.Println("El dato ingresado es incorrecto")
// 			scanner = bufio.NewScanner(os.Stdin)
// 			numero, err = strconv.Atoi(scanner.Text()) //convertimos string a int
// 		}
// 	}
// 	for i := 1; i < 11; i++ {
// 		fmt.Println(numero * i)
// 	}
// }

func TabladeMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin) //stdin es standard input (el teclado)

	for {
		fmt.Println("Ingrese un número: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, i*numero)
	}
}
