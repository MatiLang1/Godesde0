package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b

}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMidd(sumar)(5, 9)
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(5, 7)
	fmt.Println(result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	//"f" es nombre de la función (el dato que le estamos dando), los 2 primeros int son el tipo de dato de los 2 primeros n y el tercer int es porque la funcion devuelve un int
	return func(a, b int) int {
		fmt.Println("Inicio de Operación")
		return f(a, b)
	}

}

//los middlewares reciben y devuelven el mismo tipo de dato
